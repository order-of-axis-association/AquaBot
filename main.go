package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"regexp"

	"github.com/bwmarrin/discordgo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"


	"github.com/order-of-axis-association/AquaBot/types"

	"github.com/order-of-axis-association/AquaBot/db"

	"github.com/order-of-axis-association/AquaBot/admin"
	"github.com/order-of-axis-association/AquaBot/argparse"
	"github.com/order-of-axis-association/AquaBot/funcs"
	"github.com/order-of-axis-association/AquaBot/triggers"
	"github.com/order-of-axis-association/AquaBot/webhooks"
)

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
}

var global_state = types.G_State{}

var token string
var buffer = make([][]byte, 0)

func main() {

	if token == "" {
		fmt.Println("No token provided. Please run: aqua -t <bot token>")
		return
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// Register ready as a callback for the ready events.
	dg.AddHandler(ready)

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// Register guildCreate as a callback for the guildCreate events.
	dg.AddHandler(guildCreate)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	// Start github webhook listener server

	go webhooks.InitWebhookServer(dg)

	dsn := db.BuildCloudSQLDSN()
	gorm_db, err := gorm.Open("mysql", dsn)
	global_state.DBConn = gorm_db
	fmt.Println("DB err: ", err)
	fmt.Println("dbconn: %+v", global_state.DBConn)

	db.Migrate(global_state)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Aqua is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) when the bot receives
// the "ready" event from Discord.
func ready(s *discordgo.Session, event *discordgo.Ready) {
	status := os.Getenv("COMMIT_SHA")

	if status != "" {
		status = "Running on SHA: " + status
	} else {
		status = "Nani the Fuck"
	}

	fmt.Println("Setting status to: ", status)

	// Set the playing status.
	s.UpdateStatus(0, status)
}

// These funcs are meant to be "real" functionality of the bot
// where invocation requires prepending the entire message with a !
func routeMessageFunc(message string, s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("Starting route logic")

	func_maps := []map[string]interface{}{
		funcs.FuncMap,
		admin.FuncMap,
	}

	for _, func_map := range func_maps {
		for f_str, f := range func_map {

			//argparse.ParseCommandString(message)

			f_str_lower := strings.ToLower(f_str)
			f_str_lower_runes := []rune(f_str_lower)
			message_runes := []rune(message)

			fmt.Println("Attempting to route func for:", f_str, f_str_lower)

			if strings.HasPrefix(strings.ToLower(message), f_str_lower) {
				f.(func(string, *discordgo.Session, *discordgo.MessageCreate, types.G_State))(
					string(message_runes[len(f_str_lower_runes):]), // This annoying shit to preserve casing in the message
					s,
					m,
					global_state)
			}
		}
	}
}

// Autotriggers are separate from regular bot "functions".
// Autotriggers are meant to be lightweight and "fun" things that react
// to certain regexes in messages - regardless of whether there was a ! at the beginning of the command.
// Eg, one func is triggers.UselessAqua which just adds an emote to any message that has "useless" or "aqua" in it.
func routeAutoTriggers(message string, s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("Seeing if message applies to any auto-react triggers")

	var re *regexp.Regexp

	for regex, f := range triggers.FuncMap {
		re = regexp.MustCompile(regex)
		if (re.MatchString(message)) {
			f.(func(string, *discordgo.Session, *discordgo.MessageCreate, types.G_State)) (
				message,
				s,
				m,
				global_state)
		}
	}
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Imports will import new records into db so we have a corresponding match
	// These do nothing beyond ensuring a single record exists in the DB for each corresponding entity.
	db.ImportGuild(m.GuildID, global_state)
	db.ImportChannel(m.ChannelID, global_state)
	db.ImportUser(m.Author, global_state)
	db.ImportUsers(m.Mentions, global_state)

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Util commands
	if strings.HasPrefix(m.Content, "!") {
		fmt.Println("Routing util command with message:", m.Content)
		routeMessageFunc(strings.TrimLeft(m.Content, "!"), s, m)
	}

	// See if message triggers any of the autotriggers
	routeAutoTriggers(m.Content, s, m)
}

// guild is joined.
func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {

	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			_, _ = s.ChannelMessageSend(channel.ID, "Nee Nee Kazuma~")
			return
		}
	}
}

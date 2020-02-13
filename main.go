package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/order-of-axis-association/AquaBot/argparse"
	"github.com/order-of-axis-association/AquaBot/autodelete"
	"github.com/order-of-axis-association/AquaBot/cleverbot"
	"github.com/order-of-axis-association/AquaBot/config"
	"github.com/order-of-axis-association/AquaBot/db"
	"github.com/order-of-axis-association/AquaBot/triggers"
	"github.com/order-of-axis-association/AquaBot/types"
	"github.com/order-of-axis-association/AquaBot/utils"
	"github.com/order-of-axis-association/AquaBot/webhooks"
)

func init() {
	flag.StringVar(&TOKEN, "t", "", "Bot Token")
	flag.Parse()
}

var G_STATE = types.G_State{}
var CBPAYLOAD_CHAN = make(chan types.CBPayload)

var TOKEN string
var BUFFER = make([][]byte, 0)

func main() {

	if TOKEN == "" {
		fmt.Println("No token provided. Please run: aqua -t <bot token>")
		return
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + TOKEN)
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

	// Create DB connection
	dsn := db.BuildCloudSQLDSN()
	gorm_db, err := gorm.Open("mysql", dsn)
	G_STATE.DBConn = gorm_db
	fmt.Println("DB err: ", err)
	fmt.Println("dbconn: %+v", G_STATE.DBConn)

	db.Migrate(G_STATE)

	// Start github webhook listener server
	go webhooks.InitWebhookServer(dg)

	// Start auto message deleter
	go autodelete.AutoDeleter(dg, G_STATE)

	// Start Cleverbot "daemon".
	go cb.StartCBDaemon(G_STATE, CBPAYLOAD_CHAN)

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
// These will naturally not have a types.CmdArgs due to the lack of a "command" input
func routeMessageFunc(message string, state types.MessageState) {
	fmt.Println("Starting route logic for: " + message)
	for _, func_config := range config.EnabledFuncPackages {
		commands := func_config.Commands
		prefix := func_config.Prefix

		var trimmed_message string
		if !strings.HasPrefix(message, prefix) {
			continue
		} else {
			trimmed_message = strings.TrimLeft(message, prefix)
		}

		for _, command := range commands {
			cmd_str := command.Cmd
			version := command.Version
			f := command.Func

			flag_config := make(map[string]string)
			if command.Flags != nil {
				flag_config = command.Flags
			}

			cmd_args, err := argparse.ParseCommandString(trimmed_message, flag_config)
			if err != nil {
				fmt.Println("Could not parse this command. Skipping. Input was:", trimmed_message)
				utils.ApplyErrorReaction(state)
				return
			} else {
				fmt.Sprintln("%+v", cmd_args)
			}

			if strings.ToLower(cmd_args.Cmd) == strings.ToLower(cmd_str) {
				func_err := f.(func(types.CmdArgs, types.MessageState) error)(
					cmd_args,
					state)

				if func_err != nil {
					msg := fmt.Sprintf("[Error][Func: %s@ver%s] %s", prefix+cmd_args.Cmd, version, func_err.Error())
					utils.Error(msg, state)
				}
			}
		}
	}
}

// Autotriggers are separate from regular bot "functions".
// Autotriggers are meant to be lightweight and "fun" things that react
// to certain regexes in messages - regardless of whether there was a ! at the beginning of the command.
// Eg, one func is triggers.UselessAqua which just adds an emote to any message that has "useless" or "aqua" in it.
func routeAutoTriggers(message string, state types.MessageState) {
	fmt.Println("Seeing if message applies to any auto-react triggers")

	var re *regexp.Regexp

	for regex, f := range triggers.FuncMap {
		re = regexp.MustCompile(regex)
		fmt.Sprintln("message: %s - regex: %s", message, regex)
		if re.MatchString(message) {
			// This is really disgusting.
			// TODO: Need to figure out a cleaner way to pass additional chan types.CBPayload arg
			// to the cleverbot invocation function.
			// Questions: Do I prefer a consistent triggerfunc signature? Should CB be a one-off function? The invocation via regex trigger fits nicely tho.
			// For the love of god please clean me up eventually.

			var aqua_str string = "603252075006001152>"
			fmt.Sprintln("Message is %s - looking for %s", message, aqua_str)
			if strings.Contains(message, aqua_str) {
				fmt.Sprintln("Sending message, '%s' on CBPayload channel...", message)
				f.(func(string, types.MessageState, chan types.CBPayload))(
					message,
					state,
					CBPAYLOAD_CHAN,
				)
			} else {
				f.(func(string, types.MessageState))(
					message,
					state)
			}
		}
	}
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Imports will import new records into db so we have a corresponding match
	// These do nothing beyond ensuring a single record exists in the DB for each corresponding entity.
	db.ImportGuild(m.GuildID, G_STATE)
	db.ImportChannel(m.ChannelID, G_STATE)
	db.ImportUser(m.Author, G_STATE)
	db.ImportUsers(m.Mentions, G_STATE)

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	state := types.MessageState{
		S: s,
		M: m,
		G: &G_STATE,
	}

	// Command invocation
	routeMessageFunc(m.Content, state)

	// See if message triggers any of the autotriggers
	routeAutoTriggers(m.Content, state)
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

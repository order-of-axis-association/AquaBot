package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/order-of-axis-association/AquaBot/funcs"
	"github.com/order-of-axis-association/AquaBot/webhooks"
)

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
}

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

	status := os.Getenv("REPO_REVISION")

	if status != "" {
		status = "Running on SHA: " + status
	} else {
		status = "Nani the Fuck"
	}

	// Set the playing status.
	s.UpdateStatus(0, status)
}

func routeMessageFunc(message string, s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("Starting route logic")
	for f_str, f := range funcs.FuncMap {
		f_str_lower := strings.ToLower(f_str)
		f_str_lower_runes := []rune(f_str_lower)
		message_runes := []rune(message)

		fmt.Println("Attempting to route func for:", f_str, f_str_lower)

		if strings.HasPrefix(strings.ToLower(message), f_str_lower) {
			f.(func(string, *discordgo.Session, *discordgo.MessageCreate))(
				string(message_runes[len(f_str_lower_runes):]), // This annoying shit to preserve casing in the message
				s,
				m)
		}
	}
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!") {
		fmt.Println("Routing util command with message:", m.Content)
		routeMessageFunc(strings.TrimLeft(m.Content, "!"), s, m)
	}
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

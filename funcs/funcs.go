package funcs

import (
	"fmt"
	"github.com/order-of-axis-association/AquaBot/types"
	"github.com/bwmarrin/discordgo"
)

func Ping(cmd_args types.CmdArgs, s *discordgo.Session, m *discordgo.MessageCreate, global_state types.G_State) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		// Could not find channel.
		return
	}

	_, err = s.ChannelMessageSend(c.ID, "Pong")
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}

func Pong(cmd_args types.CmdArgs, s *discordgo.Session, m *discordgo.MessageCreate, global_state types.G_State) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		// Could not find channel.
		return
	}

	_, err = s.ChannelMessageSend(c.ID, "Ping")
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}

func Help(cmd_args types.CmdArgs, s *discordgo.Session, m *discordgo.MessageCreate, global_state types.G_State) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Could not find channel.")
		return
	}

	s.ChannelMessageSend(c.ID, "I'm a cute useless godess with a great ass. Leave me alone.")
}

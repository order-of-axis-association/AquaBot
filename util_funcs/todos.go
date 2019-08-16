package util_funcs

import (
	"fmt"

	"github.com/order-of-axis-association/AquaBot/types"

	"github.com/bwmarrin/discordgo"
)

func Todo(cmd_args types.CmdArgs, s *discordgo.Session, m *discordgo.MessageCreate, global_state types.G_State) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return
	}

	_, err = s.ChannelMessageSend(c.ID, "WIP!")
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}

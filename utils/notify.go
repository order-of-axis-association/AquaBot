package utils

// Functions to "notify" people with.

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Say(message string, s *discordgo.Session, m *discordgo.MessageCreate) error {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return err
	}

	_, err = s.ChannelMessageSend(c.ID, message)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return err
	}

	return nil
}

func Error(message string, s *discordgo.Session, m *discordgo.MessageCreate) error {
	message = fmt.Sprintf("❌ *%s*", message)
	return Say(message, s, m)
}

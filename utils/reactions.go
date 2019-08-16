package utils

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Must provide fallback_emoji_unicode as a unicde symbol
func ApplyReaction(emoji_name string, fallback_emoji_unicode string, s *discordgo.Session, m *discordgo.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Could not find channel by ID:", err)
		return
	}
	g, err := s.Guild(m.GuildID)
	if err != nil {
		fmt.Println("Could not find guild by ID:", err)
		return
	}

	for _, emoji := range g.Emojis {
		if emoji.Name == emoji_name {
			err = s.MessageReactionAdd(c.ID, m.ID, emoji.APIName())
			if err != nil {
				fmt.Println("Error applying emoji to message:", err)
			}
			return
		}
	}

	err = s.MessageReactionAdd(c.ID, m.ID, fallback_emoji_unicode)
	if err != nil {
		fmt.Println("Failed applying fallback emoji. Are you sure it's a default emoji?", err)
	}
}

func ApplyErrorReaction(s *discordgo.Session, m *discordgo.MessageCreate) {
	ApplyReaction("error", "⁉", s, m)
}

package utils

import (
	"fmt"
	"github.com/order-of-axis-association/AquaBot/types"
)

// Must provide fallback_emoji_unicode as a unicde symbol
func ApplyReaction(emoji_name string, fallback_emoji_unicode string, state types.MessageState) {
	m := state.M
	s := state.S
	c, err := state.S.State.Channel(m.ChannelID)
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

func ApplyErrorReaction(state types.MessageState) {
	ApplyReaction("error", "⁉", state)
}

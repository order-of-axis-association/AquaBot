package triggers

import (
	"fmt"
	"github.com/order-of-axis-association/AquaBot/types"
	"github.com/bwmarrin/discordgo"
)

func UselessAqua(args string, s *discordgo.Session, m *discordgo.MessageCreate, global_state types.G_State) {
	fmt.Println("Here")
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Could not find channel.")
		return
	}

	g, err := s.Guild(m.GuildID)
	if err != nil {
		fmt.Println("Could not find guild.")
		return
	}

	for _, emoji := range g.Emojis {
		fmt.Println(emoji.ID, emoji.Name)
		if (emoji.Name == "nani") {
			fmt.Println(emoji.APIName())
			err = s.MessageReactionAdd(c.ID, m.ID, emoji.APIName())
			fmt.Println(err)
		}
	}
}

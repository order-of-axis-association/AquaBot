package triggers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/order-of-axis-association/AquaBot/types"
	"github.com/order-of-axis-association/AquaBot/utils"
)

func UselessAqua(args string, s *discordgo.Session, m *discordgo.MessageCreate, global_state types.G_State) {
	utils.ApplyReaction("nani", "👋", s, m)
}

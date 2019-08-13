package triggers

import (
	"github.com/order-of-axis-association/AquaBot/types"
	"github.com/order-of-axis-association/AquaBot/utils"
	"github.com/bwmarrin/discordgo"
)

func UselessAqua(args string, s *discordgo.Session, m *discordgo.MessageCreate, global_state types.G_State) {
	utils.ApplyReaction("nani", "👋", s, m)
}

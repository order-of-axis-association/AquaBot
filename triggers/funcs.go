package triggers

import (
	"github.com/order-of-axis-association/AquaBot/types"
	"github.com/order-of-axis-association/AquaBot/utils"
)

func UselessAqua(args string, state types.MessageState) {
	utils.ApplyReaction("nani", "👋", state)
}

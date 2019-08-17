package types

import (
	"github.com/bwmarrin/discordgo"
)

type MessageState struct {
	S *discordgo.Session
	M *discordgo.MessageCreate
	G *G_State
}

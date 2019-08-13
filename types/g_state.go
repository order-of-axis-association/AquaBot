package types

import (
	_ "github.com/bwmarrin/discordgo"
	"github.com/jinzhu/gorm"
)

// Ugh. To add discordgo.Session to this or not...
// As it currently stands, there's only ever one session so any given discordgo.Session in
// the event handlers SHOULD be the "global" session but...
type G_State struct {
	DBConn *gorm.DB
}

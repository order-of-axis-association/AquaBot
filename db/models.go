package db

import (
	"github.com/jinzhu/gorm"
)

var StringToModelMap = map[string]interface{}{
	"guild":   Guild{},
	"channel": Channel{},
	"user":    User{},
}

// I know I really should be adding struct tags but
// GORM's documentation is terrible and it seems there's been
// a sudden uptick in activity through PRs and whatnot.
// I'm gonna hold off on adding struct tags in the hope that
// GORM will actually add things like FKs and update the struct tag docus.
// As is, it's really badly documented and some of the "features" on the docs don't actually work.

// Servers are called Guilds in Discord docs
type Guild struct {
	gorm.Model

	GuildId string
}

// ---------------------------
type Channel struct {
	gorm.Model

	ChannelId string
}

// ---------------------------
type User struct {
	gorm.Model

	UserId        string
	Username      string
	Descriminator string // The 4 numbers after the discord name
	Bot           bool
}

// ---------------------------
type Reminder struct {
	gorm.Model

	UserId   string
	Message  string
	RemindAt int // Should be an epoch
	Reminded bool
}

// ---------------------------
type Todo struct {
	gorm.Model

	UserID string
	Task   string

	Done     bool
	DoneDate int
}

// ---------------------------
type Config struct {
	gorm.Model

	ServerId      string
	ChannelId     string
	LastUserToSet int

	ConfigName  string
	ConfigValue string
}

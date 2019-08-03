package db

import (
	"github.com/jinzhu/gorm"
)

// I know I really should be adding struct tags but
// GORM's documentation is terrible and it seems there's been
// a sudden uptick in activity through PRs and whatnot.
// I'm gonna hold off on adding struct tags in the hope that
// GORM will actually add things like FKs and update the struct tag docus.
// As is, it's really badly documented and some of the "features" on the docs don't actually work.

type Server struct {
	gorm.Model

	ServerId	int
}

type Channel struct {
	gorm.Model

	ChannelId	int
}

type User struct {
	gorm.Model

	UserId		int
}

type Reminder struct {
	gorm.Model

	UserId		int
	Message		string
	RemindAt	int // Should be an epoch
}

type Todo struct {
	gorm.Model

	UserID			int
	Task			string
}

type Config struct {
	gorm.Model

	ServerId		int
	ChannelId		int
	LastUserToSet	int

	ConfigName		string
	ConfigValue		string
}

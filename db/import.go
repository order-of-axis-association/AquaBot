package db

import (
	"fmt"
	_ "strconv"

	"github.com/order-of-axis-association/AquaBot/types"

	"github.com/bwmarrin/discordgo"

	_ "github.com/jinzhu/gorm"
)

func ImportGuild(guild_id string, g_state types.G_State) {
	db := g_state.DBConn

	guild := Guild{}
	if not_found := db.Where("guild_id = ?", guild_id).First(&guild).RecordNotFound(); not_found {
		guild = Guild{GuildId: guild_id}
		db.Create(&guild)
		fmt.Println("Created new Guild record for", guild_id)
	}
}

func ImportChannel(chan_id string, g_state types.G_State) {
	db := g_state.DBConn

	channel := Channel{}
	if not_found := db.Where("channel_id = ?", chan_id).First(&channel).RecordNotFound(); not_found {
		channel = Channel{ChannelId: chan_id}
		db.Create(&channel)
		fmt.Println("Created new Channel record for", chan_id)
	}
}

func ImportUser(user_obj *discordgo.User, g_state types.G_State) {
	db := g_state.DBConn
	user_id := user_obj.ID
	username := user_obj.Username
	descriminator := user_obj.Discriminator

	user := User{}
	if not_found := db.Where("user_id = ?", user_id).First(&user).RecordNotFound(); not_found {
		user = User{UserId: user_id, Username: username, Descriminator: descriminator}
		db.Create(&user)
		fmt.Println("Created new User record for", user)
	}
}

func ImportUsers(users []*discordgo.User, g_state types.G_State) {
	for _, user := range users {
		ImportUser(user, g_state)
	}
}

package db

import (
	"fmt"
	"strconv"

	"github.com/order-of-axis-association/AquaBot/types"

	"github.com/bwmarrin/discordgo"

	_ "github.com/jinzhu/gorm"
)

func ImportGuild (guild_id_s string, g_state types.G_State) {
	db := g_state.DBConn

	guild_id, err := strconv.Atoi(guild_id_s)
	if err != nil {
		fmt.Println("Could not convert guild_id to int:", err)
	}

	server := Server{ServerId: guild_id}
	if result := db.NewRecord(server); result {
		//db.Create(&server)
		fmt.Println("Created new Server record for",guild_id)
	}
	fmt.Println("Server record already existed")

	var serv Server
	if err := db.Where("server_id = ?", guild_id).First(&serv); err != nil {
		fmt.Println("Could not find server by", guild_id)
	}

	fmt.Println("%+v", serv)

	fmt.Println("AAAAAAAAAAAA")
	var servs []Server
	db.Find(&servs)
	for _, serv := range servs {
		fmt.Println("%+v", serv)
	}
	fmt.Println("Done")

}

func ImportChannel (chann_id string, g_state types.G_State) {
	db := g_state.DBConn

	var channel Channel
	if err := db.Where("channel_id = ?", chann_id).First(&channel); err != nil {
		fmt.Println("Could not find channel by", chann_id)
	}

	fmt.Println("%+v", channel)
}

func ImportUser (user_obj *discordgo.User, g_state types.G_State) {
	db := g_state.DBConn

	user_id := user_obj.ID

	var user User
	if err := db.Where("user_id = ?", user_id).First(&user); err != nil {
		fmt.Println("Could not find user by", user_id)
	}

	fmt.Println("%+v", user)
}

func ImportUsers(users []*discordgo.User, g_state types.G_State) {
	for _, user := range users {
		ImportUser(user, g_state)
	}
}

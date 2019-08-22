package autodelete

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/order-of-axis-association/AquaBot/db"
	"github.com/order-of-axis-association/AquaBot/types"
)

func deleteMessages(s *discordgo.Session, g_state types.G_State) {
	var to_delete []db.TempMessage
	if err := g_state.DBConn.Find(&to_delete).Error; err != nil {
		fmt.Println("Err:", err)
		return
	}

	for _, msg := range to_delete {
		delete_at := msg.CreatedAt.Add((time.Second * time.Duration(msg.Length)))
		if time.Now().After(delete_at) {
			err := s.ChannelMessageDelete(msg.ChannelId, msg.MessageId)
			if err != nil {
				fmt.Sprintln("Failed to delete message")
			} else {
				g_state.DBConn.Delete(&msg)
			}
		}
	}
}

func AutoDeleter(s *discordgo.Session, g_state types.G_State) {
	for {
		time.Sleep(1 * time.Second)
		go deleteMessages(s, g_state)
	}
}

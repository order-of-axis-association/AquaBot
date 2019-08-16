package admin

import (
	"fmt"
	"errors"

	_ "github.com/order-of-axis-association/AquaBot/db"
	_ "github.com/order-of-axis-association/AquaBot/types"
	_ "github.com/order-of-axis-association/AquaBot/utils"

	"github.com/bwmarrin/discordgo"
	_ "github.com/jinzhu/gorm"
)

func isAdminUser(s *discordgo.Session, m *discordgo.MessageCreate) (bool, error) {
	valid_users := []string{
		"Remi#5619",
	}

	author_id_string := m.Author.Username + "#" + m.Author.Discriminator

	fmt.Println(author_id_string)

	for _, valid_user := range valid_users {
		if valid_user == author_id_string {
			return true, nil
		}
	}

	return false, errors.New("The user " + author_id_string + " is not an authorized admin user.")
}

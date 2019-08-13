package admin

import (
	"fmt"
	"errors"

	"github.com/order-of-axis-association/AquaBot/types"
	"github.com/order-of-axis-association/AquaBot/utils"

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


func ClearDBFlags() map[string]string {
	return map[string]string{
		"t": "table",
	}
}

func ClearDB(cmd_args types.CmdArgs, s *discordgo.Session, m *discordgo.MessageCreate, g_state types.G_State) error {
	if is_admin, err := isAdminUser(s, m); ! is_admin {
		return err
	}

	if val, ok := cmd_args.FlagArgs["table"]; ok {
		fmt.Println("Processing cleardb command...")
		return utils.Say(val, s, m)
		//g_state.DBConn.
	} else {
		fmt.Println("Not enough args")
		return utils.Error("Must provide a -t/--table to clear!", s, m)
	}

	return nil
}

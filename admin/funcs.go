package admin

import (
	"fmt"
	"errors"

	"github.com/order-of-axis-association/AquaBot/db"
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

func ClearTableFlags() map[string]string {
	return map[string]string{
		"m": "model",
	}
}

func ClearTable(cmd_args types.CmdArgs, s *discordgo.Session, m *discordgo.MessageCreate, g_state types.G_State) error {
	if is_admin, err := isAdminUser(s, m); ! is_admin {
		return err
	}

	model, ok := cmd_args.FlagArgs["model"]
	if ! ok {
		return utils.Error("Must provide a -m/--model to clear!", s, m)
	}

	model_obj, ok := db.StringToModelMap[model]
	if ! ok {
		return utils.Error("Invalid model name provided!", s, m)
	}

	if err := g_state.DBConn.Delete(model_obj).Error; err != nil {
		return utils.Error(fmt.Sprintf("Could not delete records for model '%s' Error:", err), s, m)
	}

	return utils.Say(fmt.Sprintf("Successfully deleted all records for model '%s'", strings.Title(model)), s, m)
}

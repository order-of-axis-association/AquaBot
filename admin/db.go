package admin

import (
	"fmt"
	"strings"

	"github.com/order-of-axis-association/AquaBot/db"
	"github.com/order-of-axis-association/AquaBot/utils"
	"github.com/order-of-axis-association/AquaBot/types"

	"github.com/bwmarrin/discordgo"
)

func ClearTableFlags() (map[string]string) {
	return map[string]string{
		"m": "model",
	}
}

func ClearTable(cmd_args types.CmdArgs, s *discordgo.Session, m *discordgo.MessageCreate, g_state types.G_State) (error) {
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

func CheckTables(cmd_args types.CmdArgs, s *discordgo.Session, m *discordgo.MessageCreate, g_state types.G_State) (error) {
	if is_admin, err := isAdminUser(s, m); ! is_admin {
		return err
	}

	query, ok := cmd_args.FlagArgs["query"]
	if ! ok {
		return utils.Error("Must provide a -q/--query to execute!", s, m)
	}

	err := g_state.DBConn.Exec(query).Error

	if err != nil {
		return utils.Error(fmt.Sprintf("Error executing statement: %s", query), s, m)
	}

	return nil
}

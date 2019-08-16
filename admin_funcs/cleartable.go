package admin_funcs

import (
	"fmt"
	"strings"

	"github.com/order-of-axis-association/AquaBot/db"
	"github.com/order-of-axis-association/AquaBot/types"
	"github.com/order-of-axis-association/AquaBot/utils"

	"github.com/bwmarrin/discordgo"
)


var ClearTable = types.Command {
	Cmd: "cleartable",
	Version: "0.1.0",

	Func: ClearTableFunc,
	Flags: ClearTableFlags(),
	Usage: ClearTableUsage,
}

func ClearTableFlags() types.FuncFlags {
	return types.FuncFlags{
		"m": "model",
	}
}

var ClearTableUsage = `
$cleartable
	- Shows help
$cleartable -m/--model <modelname>
	- Soft deletes all records in db for <modelname>
`

func ClearTableFunc(cmd_args types.CmdArgs, s *discordgo.Session, m *discordgo.MessageCreate, g_state types.G_State) error {
	if is_admin, err := isAdminUser(s, m); !is_admin {
		return err
	}

	model, ok := cmd_args.FlagArgs["model"]
	if !ok {
		return utils.Error("Must provide a -m/--model to clear!", s, m)
	}

	model_obj, ok := db.StringToModelMap[model]
	if !ok {
		return utils.Error("Invalid model name provided!", s, m)
	}

	if err := g_state.DBConn.Delete(model_obj).Error; err != nil {
		return utils.Error(fmt.Sprintf("Could not delete records for model '%s' Error:", err), s, m)
	}

	return utils.Say(fmt.Sprintf("Successfully deleted all records for model '%s'", strings.Title(model)), s, m)
}

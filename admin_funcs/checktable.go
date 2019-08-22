package admin_funcs

import (
	"fmt"
	_ "strings"

	_ "github.com/order-of-axis-association/AquaBot/db"
	"github.com/order-of-axis-association/AquaBot/types"
	"github.com/order-of-axis-association/AquaBot/utils"
)

var CheckTable = types.Command{
	Cmd:     "checktable",
	Version: "0.0.1",

	Func:  CheckTableFunc,
	Flags: nil,
	Usage: "",
}

func CheckTableFunc(cmd_args types.CmdArgs, state types.MessageState) error {
	if is_admin, err := utils.IsAdmin(state); !is_admin {
		return err
	}

	query, ok := cmd_args.FlagArgs["query"]
	if !ok {
		return utils.Error("Must provide a -q/--query to execute!", state)
	}

	err := state.G.DBConn.Exec(query).Error

	if err != nil {
		return utils.Error(fmt.Sprintf("Error executing statement: %s", query), state)
	}

	return nil
}

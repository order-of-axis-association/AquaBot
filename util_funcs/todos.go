package util_funcs

import (
	"errors"
	"fmt"

	"github.com/order-of-axis-association/AquaBot/types"
	"github.com/order-of-axis-association/AquaBot/utils"

	"github.com/bwmarrin/discordgo"
)

var Todo = types.Command {
	Cmd: "todo",
	Version: "0.0.1",

	Func: TodoFunc,
	Flags: nil,
	Usage: TodoUsage,
}

var TodoUsage = `
!todo
	- Lists all current todos.
!todo list
	- Same as above
!todo add <msg>
	- Add <msg> to your todo list
!todo rem <numerical_id>
	- Given the todo ID given in "!todo list", will delete corresponding entry.
`

func TodoFunc(cmd_args types.CmdArgs, s *discordgo.Session, m *discordgo.MessageCreate, global_state types.G_State) error {
	utils.Say("WIP!", s, m)

	pos_args := cmd_args.PosArgs
	if len(pos_args) == 0 {
		msg := "Not enough args supplied!"
		utils.Error(msg, s, m);
		return errors.New(msg)
	}
	first_arg := cmd_args.PosArgs[0]

	fmt.Println(first_arg)
	return nil
}

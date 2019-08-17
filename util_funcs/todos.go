package util_funcs

import (
	_ "errors"
	"strings"
	"regexp"
	"strconv"
	"fmt"

	"github.com/order-of-axis-association/AquaBot/db"
	"github.com/order-of-axis-association/AquaBot/types"
	"github.com/order-of-axis-association/AquaBot/utils"

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
	- Given the todo ID given in '!todo list', will delete corresponding entry.
	  Eg, if you have three todos listed as 1), 2) and 3) and
	  you want to delete the middle todo, one would use
	  '!todo rem 2'
`

func TodoFunc(cmd_args types.CmdArgs, state types.MessageState) error {
	pos_args := cmd_args.PosArgs

	if len(pos_args) == 0 {
		return displayAllTodos(state)
	}

	first_arg := pos_args[0]
	remaining_args := strings.Join(pos_args[1:], " ")

	if first_arg == "list" {
		return displayAllTodos(state)
	} else if utils.StrContains(first_arg, []string{"add", "create"}) {
		return addTodo(remaining_args, state)
	} else if utils.StrContains(first_arg, []string{"rem", "remove", "del", "delete", "finish"}) {
		return remTodo(remaining_args, state)
	} else {
		return utils.Error("Invalid subcommand! Please see `!help todo` for help.", state)
	}
}

func displayAllTodos(state types.MessageState) error {
	user_id := state.M.Author.ID

	var todos []db.Todo
	state.G.DBConn.Where("user_id = ?", user_id).Find(&todos)

	if len(todos) == 0 {
		msg := "You have no todos!\nGet started with `!todo add <task>`!"
		return utils.TempSay(msg, state)
	}

	msg := "Your todos are...\n\n"

	for index, todo := range todos {
		msg += fmt.Sprintf("**%d)** %s - _Added on %s_\n", index + 1, todo.Task, todo.CreatedAt)
	}

	utils.TempSay(msg, state)

	return nil
}


func addTodo(task string, state types.MessageState) error {
	if task == "" {
		return utils.Error("You must provide text to save as your todo! `!todo add <task>`", state)
	}

	todo := db.Todo{
		UserID: state.M.Author.ID,
		Task: task,
		Done: false,
		DoneDate: -1,
	}

	if err := state.G.DBConn.Create(&todo).Error; err != nil {
		return utils.Error("Encountered error while trying to save a new Todo. Err:"+err.Error(), state)
	}
	return utils.TempSay("Successfully saved new Todo!", state)
}

func remTodo(deletion_id_str string, state types.MessageState) error {
	re := regexp.MustCompile(`^\d+$`)
	if !re.MatchString(deletion_id_str) {
		return utils.Error("You must provide a numerical id to delete! `!todo rem <num>`", state)
	}

	deletion_id, err := strconv.Atoi(deletion_id_str)
	if err != nil {
		return utils.Error("Could not convert deletion id to integer! Yell at Remi", state)
	}

	var todos []db.Todo
	state.G.DBConn.Where("user_id = ?", state.M.Author.ID).Find(&todos)

	if len(todos) < deletion_id {
		msg := fmt.Sprintf("You only have %d todo(s)! Can't delete todo number %d!", len(todos), deletion_id)
		return utils.Error(msg, state)
	}

	if err = state.G.DBConn.Delete(todos[deletion_id-1]).Error; err != nil {
		return utils.Error(fmt.Sprintf("Error while deleting Todo #%d!", deletion_id), state)
	}
	return utils.TempSay(fmt.Sprintf("Successfully deleted todo entry #%d!", deletion_id), state)
}

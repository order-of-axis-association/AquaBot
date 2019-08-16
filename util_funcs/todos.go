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
	- Given the todo ID given in '!todo list', will delete corresponding entry.
	  Eg, if you have three todos listed as 1), 2) and 3) and
	  you want to delete the middle todo, one would use
	  '!todo rem 2'
`

func TodoFunc(cmd_args types.CmdArgs, s *discordgo.Session, m *discordgo.MessageCreate, global_state types.G_State) error {
	pos_args := cmd_args.PosArgs

	if len(pos_args) == 0 {
		return displayAllTodos(s, m, global_state)
	}

	first_arg := pos_args[0]
	remaining_args := strings.Join(pos_args[1:], " ")

	if first_arg == "list" {
		return displayAllTodos(s, m, global_state)
	} else if utils.StrContains(first_arg, []string{"add", "create"}) {
		return addTodo(remaining_args, s, m, global_state)
	} else if utils.StrContains(first_arg, []string{"rem", "remove", "del", "delete", "finish"}) {
		return remTodo(remaining_args, s, m, global_state)
	} else {
		return utils.Error("Invalid subcommand! Please see `!help todo` for help.", s, m)
	}
}

func displayAllTodos(s *discordgo.Session, m *discordgo.MessageCreate, global_state types.G_State) error {
	user_id := m.Author.ID

	var todos []db.Todo
	global_state.DBConn.Where("user_id = ?", user_id).Find(&todos)

	if len(todos) == 0 {
		msg := "You have no todos!\nGet started with `!todo add <task>`!"
		return utils.Say(msg, s, m)
	}

	msg := "Your todos are...\n\n"

	for index, todo := range todos {
		msg += fmt.Sprintf("**%d)** %s - _Added on %s_\n", index + 1, todo.Task, todo.CreatedAt)
	}

	utils.Say(msg, s, m)

	return nil
}


func addTodo(task string, s *discordgo.Session, m *discordgo.MessageCreate, global_state types.G_State) error {
	if task == "" {
		return utils.Error("You must provide text to save as your todo! `!todo add <task>`", s, m)
	}

	todo := db.Todo{
		UserID: m.Author.ID,
		Task: task,
		Done: false,
		DoneDate: -1,
	}

	if err := global_state.DBConn.Create(&todo).Error; err != nil {
		return utils.Error("Encountered error while trying to save a new Todo. Err:"+err.Error(), s, m)
	}
	return utils.Say("Successfully saved new Todo!", s, m)
}

func remTodo(deletion_id_str string, s *discordgo.Session, m *discordgo.MessageCreate, global_state types.G_State) error {
	re := regexp.MustCompile(`^\d+$`)
	if !re.MatchString(deletion_id_str) {
		return utils.Error("You must provide a numerical id to delete! `!todo rem <num>`", s, m)
	}

	deletion_id, err := strconv.Atoi(deletion_id_str)
	if err != nil {
		return utils.Error("Could not convert deletion id to integer! Yell at Remi", s, m)
	}

	var todos []db.Todo
	global_state.DBConn.Where("user_id = ?", m.Author.ID).Find(&todos)

	if len(todos) < deletion_id {
		msg := fmt.Sprintf("You only have %d todo(s)! Can't delete todo number %d!", len(todos), deletion_id)
		return utils.Error(msg, s, m)
	}

	if err = global_state.DBConn.Delete(todos[deletion_id-1]).Error; err != nil {
		return utils.Error(fmt.Sprintf("Error while deleting Todo #%d!", deletion_id), s, m)
	}
	return utils.Say(fmt.Sprintf("Successfully deleted todo entry #%d!", deletion_id), s, m)
}

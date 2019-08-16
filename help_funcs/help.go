package help_funcs
// Help is its own package due to circular import dependency issues if we included it in util_funcs

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/order-of-axis-association/AquaBot/util_funcs"
	"github.com/order-of-axis-association/AquaBot/types"
	"github.com/order-of-axis-association/AquaBot/utils"
)

var Help = types.Command {
	Cmd: "help",
	Version: "0.0.1",

	Func: HelpFunc,
	Flags: nil,
	Usage: HelpUsage,
}

var HelpUsage = `
help
	- Display usage information for this command.
help <cmdname>
	- Displays help information for <cmdname>
`

func HelpFunc(cmd_args types.CmdArgs, s *discordgo.Session, m *discordgo.MessageCreate, global_state types.G_State) error {
	pos_args := cmd_args.PosArgs

	if len(pos_args) == 0 {
		// !help w/ no args
		return utils.Mono(addFuncPrefix(HelpUsage, package_prefix, cmd_args.Cmd), s, m)
	}

	known_help_funcs := getAllActiveCommands()
	help_func := pos_args[0]

	if !utils.StrContains(help_func, known_help_funcs) {
		known_help_funcs_exploded := strings.Join(known_help_funcs, ", ")
		msg := fmt.Sprintf("I don't know about the `%s` command!\nKnown commands: `%s`", help_func, known_help_funcs_exploded)
		return utils.Error(msg, s, m)
	}

	return utils.Say("I'm a cute useless godess with a great ass. Leave me alone.", s, m)
}

func getAllActiveCommands() []string {
	var funcs = make([]string, 0)

	enabled_funcs := []types.FuncPackageConfig {
		util_funcs.Config,
	}

	for _, func_config := range enabled_funcs {
		for _, command := range func_config.Commands {
			funcs = append(funcs, command.Cmd)
		}
	}

	sort.Strings(funcs)
	return funcs
}

func addFuncPrefix(usage_info string, prefix string, cmdname string) string {
	var re *regexp.Regexp
	var new_usage string

	re = regexp.MustCompile("^" + cmdname)
	new_usage = re.ReplaceAllString(usage_info, prefix + cmdname)

	re = regexp.MustCompile("\n" + cmdname)
	new_usage = re.ReplaceAllString(usage_info, "\n" + prefix + cmdname)

	return new_usage
}

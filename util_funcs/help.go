package util_funcs

import (
	_ "fmt"
	"regexp"

	"github.com/bwmarrin/discordgo"
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
		return utils.Mono(addFuncPrefix(HelpUsage, package_prefix, cmd_args.Cmd), s, m)
	} else {
		return utils.Say("I'm a cute useless godess with a great ass. Leave me alone.", s, m)
	}

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

package help_funcs

// Help is its own package due to circular import dependency issues if we included it in util_funcs

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/order-of-axis-association/AquaBot/admin_funcs"
	"github.com/order-of-axis-association/AquaBot/config_funcs"
	"github.com/order-of-axis-association/AquaBot/types"
	"github.com/order-of-axis-association/AquaBot/util_funcs"
	"github.com/order-of-axis-association/AquaBot/utils"
)

var Help = types.Command{
	Cmd:     "help",
	Version: "0.0.1",

	Func:  HelpFunc,
	Flags: nil,
	Usage: HelpUsage,
}

var HelpUsage = `
help
	- Display usage information for this command.

help <cmdname>
	- Displays help information for <cmdname>
`

var ENABLED_PACKAGES = []types.FuncPackageConfig{
	config_funcs.Config,
	util_funcs.Config,
}

var ADMIN_PACKAGES = []types.FuncPackageConfig{
	admin_funcs.Config,
}

func HelpFunc(cmd_args types.CmdArgs, state types.MessageState) error {
	pos_args := cmd_args.PosArgs
	is_admin, _ := utils.IsAdmin(state)

	if len(pos_args) == 0 {
		// !help w/ no args
		return utils.Mono(addFuncPrefix(HelpUsage, package_prefix, cmd_args.Cmd), state)
	}

	known_help_funcs, known_help_funcs_prefixed := getAllActiveCommands(is_admin)
	help_func := pos_args[0]

	if !utils.StrContains(help_func, known_help_funcs) &&
		!utils.StrContains(help_func, known_help_funcs_prefixed) {
		known_help_funcs_exploded := strings.Join(known_help_funcs_prefixed, ", ")
		msg := fmt.Sprintf("I don't know about the `%s` command!\nKnown commands: `%s`", help_func, known_help_funcs_exploded)
		return utils.Error(msg, state)
	}

	for _, pkg := range getEnabledPackages(is_admin) {
		for _, command := range pkg.Commands {
			if help_func == command.Cmd ||
				help_func == pkg.Prefix+command.Cmd {
				if command.Usage != "" {
					return utils.Mono(addFuncPrefix(command.Usage, pkg.Prefix, command.Cmd), state)
				} else {
					fmt.Println("Eh??")
					fmt.Println(command.Usage)
					return utils.TempSay("I'm a cute useless godess with a great ass. Leave me alone.", state)
				}
			}
		}
	}

	return utils.TempSay("I'm a cute useless godess with a great ass. Leave me alone.", state)
}

func getEnabledPackages(is_admin bool) []types.FuncPackageConfig {
	enabled_pkgs := ENABLED_PACKAGES
	if is_admin {
		enabled_pkgs = append(enabled_pkgs, ADMIN_PACKAGES...)
	}
	return enabled_pkgs
}

func getAllActiveCommands(is_admin bool) ([]string, []string) {
	var cmds = make([]string, 0)
	var cmds_prefixed = make([]string, 0)

	for _, pkg_config := range getEnabledPackages(is_admin) {
		for _, command := range pkg_config.Commands {
			cmds = append(cmds, command.Cmd)
			cmds_prefixed = append(cmds_prefixed, pkg_config.Prefix+command.Cmd)
		}
	}

	sort.Strings(cmds)
	sort.Strings(cmds_prefixed)
	return cmds, cmds_prefixed
}

func addFuncPrefix(usage_info string, prefix string, cmdname string) string {
	var re *regexp.Regexp
	var new_usage string

	re = regexp.MustCompile("^" + cmdname)
	new_usage = re.ReplaceAllString(usage_info, prefix+cmdname)

	re = regexp.MustCompile("\n" + cmdname)
	new_usage = re.ReplaceAllString(usage_info, "\n"+prefix+cmdname)

	return new_usage
}

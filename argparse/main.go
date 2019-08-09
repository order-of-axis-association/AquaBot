package argparse

import (
	"fmt"
	"strings"
)

// Will parse out a command from discord. Format should be:
// <cmd> [[pos_arg]..., [-<short_opt> <val>|--<long_opt> <val]...]
//
// Parses out types.CmdArgs
// Must provide all positional args BEFORE args with flags.
// Eg,
// mycmd arg1 arg2 arg3 --myflag somevalue <= VALID
// mycmd --myflag somevalue				   <= VALID
// mycmd arg1							   <= VALID
// mycmd arg1 arg2 --myflag somevalue arg3 <= INVALID
func ParseCommandString(cmd string) {
	cmd_runes := []rune(cmd)

	split_i := strings.Index(cmd, " ")
	cmdname := string(cmd_runes[:split_i])
	cmd_rem := cmd_runes[split_i+1:] // +1 to remove the space at beginning.

	fmt.Println(cmdname)
	fmt.Println(string(cmd_rem))
}

package argparse

import (
	"fmt"
	"strings"
	"regexp"
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
	// NOTE: This func is gonna get fucky if you give unicode args.
	// Should probably account for those.

	cmd_runes := []rune(cmd)

	var cmd_name string
	var cmd_rem []rune
	if strings.Contains(cmd, " ") {
		split_i := strings.Index(cmd, " ")
		cmd_name = string(cmd_runes[:split_i])
		cmd_rem = cmd_runes[split_i:] // Don't +1 the index. We want to keep the space.
	} else {
		// No spaces found, ie can't possibly have an arg after the commandname
		cmd_name = cmd
		cmd_rem = make([]rune, 0)
	}

	fmt.Println("Command name:",cmd_name)

	if len(cmd_rem) > 0 {
		// If we had args to parse...

		var positional_args, flagged_args []rune
		if strings.Contains(string(cmd_rem), " -") {
			// We have flagged args - Split out the positional args and flagged args
			first_flag_arg_i := strings.Index(string(cmd_rem), " -")

			positional_args = cmd_rem[:first_flag_arg_i]
			flagged_args = cmd_rem[first_flag_arg_i:]
		} else {
			// No flagged args.

			positional_args = cmd_rem
			flagged_args = make([]rune, 0)
		}

		fmt.Println("Positional args:", string(positional_args))
		fmt.Println("Flagged args:", string(flagged_args))

		//parsed_positional_args := 
		parsePositionalArgs(positional_args)
		//parsed_flagged_args := 
		parseFlaggedArgs(flagged_args)
	}

}

func parsePositionalArgs(cmd_rem []rune) {
	re := regexp.MustCompile(`(?P<arg>[^'" ]+|'[^']+'|"[^"]+")`)
	results := re.FindAllString(string(cmd_rem), -1)

	for _, result := range results {
		fmt.Println(result)
	}
}

func parseFlaggedArgs(cmd_rem []rune) {

}

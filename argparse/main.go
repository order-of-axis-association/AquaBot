package argparse

import (
	"fmt"
	"strings"
	"regexp"

	"github.com/order-of-axis-association/AquaBot/types"
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
func ParseCommandString(cmd string) types.CmdArgs {
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

	cmd_args := types.CmdArgs{
		Cmd: cmd_name,
	}

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

		parsed_positional_args := parsePositionalArgs(positional_args)
		parsed_flagged_args := parseFlaggedArgs(flagged_args)
		short_flagged_args := parsed_flagged_args["short"]
		long_flagged_args := parsed_flagged_args["long"]

		fmt.Println("%+v", parsed_positional_args)
		fmt.Println("%+v", short_flagged_args)
		fmt.Println("%+v", long_flagged_args)

		cmd_args.PosArgs = parsed_positional_args
		cmd_args.ShortFlagArgs = short_flagged_args
		cmd_args.LongFlagArgs = long_flagged_args
	}

	return cmd_args
}

func parsePositionalArgs(cmd_rem []rune) map[int]string {
	re := regexp.MustCompile(`(?P<arg>[^'" ]+|'[^']+'|"[^"]+")`)
	results := re.FindAllString(string(cmd_rem), -1)

	pos_arg_map := make(map[int]string)
	for index, result := range results {
		pos_arg_map[index] = result
	}

	return pos_arg_map
}

// Flagged args must be in one of the following formats:
// - `-<short_arg> <value>`
// - `-<short_arg>=<value>`
// - `--<long_arg>=<value>` <= NOTE: Long args MUST have an equal sign.
// Where:
// - <short_arg> MUST be one character from a-z, case insensitive.
// - <long_arg> MUST only contain a-z and hyphens, case insensitive. NO NUMBERS.
// - <value> can be either:
// - - Non-quoted string with no spaces
// - - Single OR double quoted string, spaces allowed.
// - - - Quoted strings can contain the "other" type of quote inside the quote. Eg, "string with ' single quote"
func parseFlaggedArgs(cmd_rem []rune) map[string]map[string]string {
	fmt.Println("-------")
	fmt.Println("======= Attempting to parse flagged args")
	fmt.Println("-------")
	// Oof lmao. Maybe I should split this into regular code and not regex.
	re := regexp.MustCompile(`(?:(?P<shortarg>-[a-zA-Z])(?:[ =](?P<shortvalue>[^- '"]+|'[^']+'|"[^"]+")?)|(?P<longarg>--[a-zA-Z-]+)=(?P<longvalue>[^- '"]+|'[^']+'|"[^"]+"))`)

	short_arg_map := make(map[string]string)
	long_arg_map := make(map[string]string)

	tmp_rem := cmd_rem
	for ok := true; ok; ok = len(tmp_rem) > 0 {
		result := re.FindSubmatchIndex([]byte(string(tmp_rem)))
		if len(result) == 0{
			fmt.Println("The remaining string was not valid. Got up to:", string(tmp_rem))
			break
		}
		fmt.Println("Looking at:", string(tmp_rem))
		//fmt.Println(result)
		//fmt.Println(re.SubexpNames())

		var k_start, k_end, v_start, v_end int
		var is_long_arg bool

		// 2nd index is start of the "shortarg" capgroup.
		// If this field is -1, that means we have a longarg.
		// Eg, [1 16 -1 -1 -1 -1 1 10 10 16]
		// Is the return for parsing the string `--flagged=value`
		// Eg, [1 17 1 3 4 17 -1 -1 -1 -1]
		// Is the return for parsing: `-h "short value"`
		if result[2] == -1 {
			// Long arg
			k_start = result[6]
			k_end = result[7]
			v_start = result[8]
			v_end = result[9]

			is_long_arg = true
		} else if result[6] == -1 {
			// Short arg
			k_start = result[2]
			k_end = result[3]
			v_start = result[4]
			v_end = result[5]

			is_long_arg = false
		}
		//fmt.Println("Accessing... k_start:", k_start, "k_end:", k_end, "v_start:", v_start, "v_end:", v_end)

		var arg_name, arg_val []rune

		arg_name = tmp_rem[k_start:k_end]
		if v_start != -1 {
			// This fun block because short args can not have a value, in which case default to empty string as the "value"
			arg_val = tmp_rem[v_start:v_end]
		} else {
			arg_val = make([]rune, 0)
		}

		fmt.Println("Parsed key:",string(arg_name), "val:", string(arg_val))

		trimmed_key_name := strings.Trim(string(arg_name), " -") // Trim spaces and dashes from key name
		trimmed_value := strings.Trim(string(arg_val), " ")

		if is_long_arg {
			long_arg_map[trimmed_key_name] = trimmed_value
		} else {
			short_arg_map[trimmed_key_name] = trimmed_value
		}

		if v_end != -1 {
			tmp_rem = tmp_rem[v_end:]
		} else {
			// Likewise, if short arg has no value, use the end of the key instead of val to cut
			tmp_rem = tmp_rem[k_end:]
		}


		fmt.Println("Remaining tmp_rem:", string(tmp_rem))
	}

	fmt.Println("++++ That's all, folks!")

	return map[string]map[string]string{
		"short": short_arg_map,
		"long": long_arg_map,
	}
}

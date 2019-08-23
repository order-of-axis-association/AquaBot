package argparse

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/order-of-axis-association/AquaBot/types"
)

// Will parse out a command from discord. Format should be:
// <cmd> [[pos_arg]..., [-<short_opt> <val>|--<long_opt> <val]...]
//
// Parses out types.CmdArgs
// Must provide all positional args BEFORE args with flags.
// Eg,
// mycmd arg1 arg2 arg3 --myflag=somevalue	<= VALID
// mycmd --myflag somevalue					<= VALID
// mycmd arg1								<= VALID
// mycmd arg1 arg2 --myflag=somevalue arg3	<= INVALID
//
// NOTE: If no FlagConfig is found for a command, all args will be DISCARDED (Original message can be found in CmdArgs.OrigMsg
func ParseCommandString(cmd string, flag_config map[string]string) (types.CmdArgs, error) {
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
		Cmd:     cmd_name,
		OrigMsg: cmd,
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

		parsed_positional_args := parsePositionalArgs(positional_args)
		short_flagged_args, long_flagged_args, err := parseFlaggedArgs(flagged_args)
		if err != nil {
			msg := fmt.Sprintln("Could not parse flagged args:", err)
			fmt.Println(msg)
			return types.CmdArgs{}, errors.New(msg)
		}

		cmd_args.PosArgs = parsed_positional_args

		if len(flag_config) > 0 {
			// If we have a flag config, let's unify the short/long args.
			cmd_args.FlagArgs, err = processFlagConfig(short_flagged_args, long_flagged_args, flag_config)
			if err != nil {
				msg := fmt.Sprintln("Could not unify flags with config:", err)
				fmt.Println(msg)
				return types.CmdArgs{}, errors.New(msg)
			}
		}
	}

	return cmd_args, nil
}

func trimEncasingQuotes(val string) string {
	valid_quotes := []string{
		"'",
		"\"",
		"`",
	}

	for _, quote := range valid_quotes {
		if string(val[0]) == quote {
			return strings.Trim(val, quote)
		}
	}

	return val
}

// This function takes the cmd_args.ShortFlagArgs and cmd_argsLongFlagArgs
// and using the input config, will map both short and long vals to FlagArgs
// The input config should be a map of short flag to long flag, eg
// [ "f": "full", "h": "help", "c": "commit" ], etc etc
// This also means this function will throw an error if both the short and long
// versions of the flag are set, eg if the command had "!cmd -c 'abc12345' --commit '12345abc'"
// This function will also trim quotes if value is encased in them.
func processFlagConfig(short_flags map[string]string, long_flags map[string]string, config map[string]string) (map[string]string, error) {
	flag_args := make(map[string]string)

	for short_flag, long_flag := range config {
		if short_val, exists := short_flags[short_flag]; exists {
			short_val = trimEncasingQuotes(short_val)
			flag_args[short_flag] = short_val
			flag_args[long_flag] = short_val
		}
		if long_val, exists := long_flags[long_flag]; exists {
			if val, exists := flag_args[short_flag]; exists {
				// If we already saw a short_flag, then we would have set both short and long flag keys, so just check one of them.
				msg := fmt.Sprintf("Both a short and long flag were set for the same flag! Short:", short_flag, "short_val:", val, "Long:", long_flag, "long_val:", long_val)
				fmt.Println(msg)
				return make(map[string]string), errors.New(msg)
			}

			long_val = trimEncasingQuotes(long_val)
			flag_args[short_flag] = long_val
			flag_args[long_flag] = long_val
		}
	}

	return flag_args, nil
}

func parsePositionalArgs(cmd_rem []rune) []string {
	re := regexp.MustCompile(`(?P<arg>[^'" ]+|'[^']+'|"[^"]+")`)
	results := re.FindAllString(string(cmd_rem), -1)

	pos_args := make([]string, 0)
	for _, result := range results {
		pos_args = append(pos_args, result)
	}

	return pos_args
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
func parseFlaggedArgs(cmd_rem []rune) (map[string]string, map[string]string, error) {
	// Oof lmao. Maybe I should split this into regular code and not regex.
	re := regexp.MustCompile(`(?:(?P<shortarg>-[a-zA-Z])|(?P<longarg>--[a-zA-Z-]+))(?:[ =](?P<value>[^- \x60'"]+|'[^']+'|"[^"]+"|\x60[^\x60]+\x60)?)`)

	short_arg_map := make(map[string]string)
	long_arg_map := make(map[string]string)

	tmp_rem := cmd_rem
	for len(tmp_rem) > 0 {
		result := re.FindSubmatchIndex([]byte(string(tmp_rem)))
		if len(result) == 0 {
			msg := fmt.Sprintln("The remaining string was not valid. Got up to:", string(tmp_rem))
			fmt.Println(msg)
			return nil, nil, errors.New(msg)
		}

		var k_start, k_end, v_start, v_end int
		var is_long_arg bool

		// 2nd index is start of the "shortarg" capgroup.
		// If this field is -1, that means we have a longarg.
		// Eg, [1 16 -1 -1 -1 -1 1 10]
		// Is the return for parsing the string ` --flagged=value`
		// Eg, [1 17 1 3 4 17 -1 -1]
		// Is the return for parsing: ` -h "short value"`
		if result[2] == -1 {
			// Long arg
			k_start = result[4]
			k_end = result[5]
			is_long_arg = true
		} else if result[6] == -1 {
			// Short arg
			k_start = result[2]
			k_end = result[3]
			is_long_arg = false
		}
		v_start = result[6]
		v_end = result[7]

		var arg_name, arg_val []rune

		arg_name = tmp_rem[k_start:k_end]
		if v_start != -1 {
			// Because args can not have a value, in which case default to empty string as the "value"
			arg_val = tmp_rem[v_start:v_end]
		} else {
			arg_val = make([]rune, 0)
		}

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
			// Likewise, if arg has no value, use the end of the key instead of val to cut
			tmp_rem = tmp_rem[k_end:]
		}
	}

	return short_arg_map, long_arg_map, nil
}

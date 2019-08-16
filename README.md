# AquaBot

As useless as its namesake.

### Adding Commands
Commands are defined by a struct, `types.Command`, which contains:
- `Cmd`, The package prefix combined with this string will invoke the command. Ie, "!help", "!todo"
- `Version`, Arbitrary version string. Mostly for dev purposes.
- `Func`, A pointer to the literal function to be executed. These funcs should have a function signature matching  
        `func MyCmd(cmd_args types.CmdArgs, s *discordgo.Session, m *discordgo.MessageCreate, global_state types.G_State) error`
- `Flags`, The `map[string]string` that maps short args to long args. If `nil` is provided here, the command will *NOT PARSE* any flagged arguments provided to this command.
- `Usage`, The "help string" that will be shown in `!help` or incorrect invocations of this command.

The set of processed, or "active", commands is defined by the commands listed in each function package (`util_funcs`, `admin_funcs`, etc) in the exposed `Config` variable (type `types.FuncPackageConfig`). This variable contains a field `Commands` of type `[]types.Commands`.

To add a new function package, update the `config/enabled_func_packages.go` file.

#### Making new Functions
All functions must follow the function signature
`func MyCmd(cmd_args types.CmdArgs, s *discordgo.Session, m *discordgo.MessageCreate, global_state types.G_State) error`

`types.CmdArgs` will contain the parsed input command in the following exposed fields:
- `Cmd`, The commandname. This matches the `Cmd` in `types.Command`.
- `OrigMsg`, The unmodified original input pre-parsing.`
- `PosArgs`, Contains "positional args". These are args provided without any flags or "keys" to match them against.  
            Eg, `!mycmd arg1 arg2 arg3` would have `[ arg1 arg2 arg3 ]` as the value for `PosArgs`.
- `FlagArgs`, Contains "flagged args". These are args that had flags, short or long. These will ONLY be parsed out if the `types.Command` config contained a valid `Flags` value. Ie, you need to tell the parser how to parse the flags.  
            Eg, if you had a `Flags` config of `{ m: myarg, s: silent, v:verbose }` and used command  
                `!mycmd -m=myvalue --silent --verbose 3`  
                You would end up with the following in `FlagArgs`:  
                `{ m:myvalue, myarg:myvalue, s:, silent:, v:3, verbose:3 }`  
                In other words, the parser adds both short and long flags to `FlagArgs` with corresponding values. If no value was provided, the value is an empty string.



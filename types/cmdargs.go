package types

type CmdArgs struct {
	Cmd				string				// Name of the command, ie "ping", "help", etc
	OrigMsg			string				// The pre-parsed original command
	PosArgs			map[int]string		// Key is positional int
	FlagArgs		map[string]string	// Calling loadCommandConfig() will populate this field
}

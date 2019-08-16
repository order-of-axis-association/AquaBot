package types

type CmdArgs struct {
	Cmd      string            // Name of the command, ie "ping", "help", etc
	OrigMsg  string            // The pre-parsed original command
	PosArgs  []string          // Array of non-flag args in order
	FlagArgs map[string]string // Calling loadCommandConfig() will populate this field
}

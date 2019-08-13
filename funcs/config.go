package funcs

import (
	"github.com/order-of-axis-association/AquaBot/types"
)

// The prefix that must be attached to any command in the FuncMap
var Prefix = "!"

// Map from the command string to the function containing its logic
var FuncMap = types.FuncMap{
	"ping": Ping,
	"pong": Pong,
	"help": Help,
	"todo": Todo,
}

// Flag maps are maps from the command name to the func that returns the flag config
// Flag configs should return a map[string]string where key is short flag name, val is long flag name.
var FlagMap = types.FlagMap{
	"ping": nil,
	"pong": nil,
	"help": nil, // <= eventually might want to add one
	"todo": nil, // <= here too
}

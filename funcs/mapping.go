package funcs

import (
	"github.com/order-of-axis-association/AquaBot/types"
)

// I guess cap camelcase cuz global?? I use underscores everywhere else :-/
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

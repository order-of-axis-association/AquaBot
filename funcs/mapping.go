package funcs

// I guess cap camelcase cuz global?? I use underscores everywhere else :-/
var FuncMap = map[string]interface{}{
	"ping": Ping,
	"pong": Pong,
	"help": Help,
	"todo": Todo,
}

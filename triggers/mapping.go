package triggers

// I guess cap camelcase cuz global?? I use underscores everywhere else :-/
var FuncMap = map[string]interface{}{
	`(?i)(aqua|useless)`:        UselessAqua, // (?i) is case insensitivity
	`(?i)<@\!?603252075006001152>`: InvokeCleverbot,
}

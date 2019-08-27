package config_cmds

var VALID_CONFIGURATORS map[string]bool = map[string]bool{
	"Remi#5619": true,
}

func validate_allowedusers(username string) bool {
	_, ok := VALID_CONFIGURATORS[username]
	return ok
}

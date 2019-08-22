package utils

import (
	"errors"
	"fmt"
	_ "strings"

	"github.com/order-of-axis-association/AquaBot/types"
)

func StrContains(needle string, haystack []string) bool {
	for _, elem := range haystack {
		if elem == needle {
			return true
		}
	}
	return false
}

func IsAdmin(state types.MessageState) (bool, error) {
	valid_users := []string{
		"Remi#5619",
	}

	author_id_string := state.M.Author.String()

	fmt.Println(author_id_string)

	for _, valid_user := range valid_users {
		if valid_user == author_id_string {
			return true, nil
		}
	}

	return false, errors.New("The user " + author_id_string + " is not an authorized admin user.")
}

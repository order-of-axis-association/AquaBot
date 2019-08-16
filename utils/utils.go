package utils

import (
	"errors"
	"fmt"
	_ "strings"

	"github.com/bwmarrin/discordgo"
)

func StrContains(needle string, haystack []string) bool {
	for _, elem := range haystack {
		if elem == needle {
			return true
		}
	}
	return false
}

func IsAdmin(s *discordgo.Session, m *discordgo.MessageCreate) (bool, error) {
	valid_users := []string{
		"Remi#5619",
	}

	author_id_string := m.Author.Username + "#" + m.Author.Discriminator

	fmt.Println(author_id_string)

	for _, valid_user := range valid_users {
		if valid_user == author_id_string {
			return true, nil
		}
	}

	return false, errors.New("The user " + author_id_string + " is not an authorized admin user.")
}

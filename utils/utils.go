package utils

import (
	_ "strings"
)

func StrContains(needle string, haystack []string) bool {
	for _, elem := range haystack {
		if elem == needle {
			return true
		}
	}
	return false
}

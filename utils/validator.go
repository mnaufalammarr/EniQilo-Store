package utils

import "regexp"

func ValidatePhoneStartsWithPlus(phone string) bool {
	match, err := regexp.MatchString(`^\+`, phone) // Matches strings starting with "+"
	return err == nil && match
}

package utils

import (
	"fmt"
	"regexp"
)

func ValidatePhoneStartsWithPlus(phone string) bool {
	match, err := regexp.MatchString(`^\+`, phone)
	if err != nil {
		fmt.Errorf("error validating phone number: %w", err)
		return false
	}
	if !match {
		fmt.Errorf("invalid phone number format (must start with + and valid international code)")
		return false
	}
	return true
}

func ValidateUrl(url string) bool {
	match, err := regexp.MatchString(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`, url)
	if err != nil {
		fmt.Errorf("error validating phone number: %w", err)
		return false
	}
	if !match {
		fmt.Errorf("invalid phone number format (must start with + and valid international code)")
		return false
	}
	return true
}

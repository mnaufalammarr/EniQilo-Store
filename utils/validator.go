package utils

import (
	"fmt"
	"regexp"
)

func ValidatePhoneStartsWithPlus(phone string) bool {
	match, err := regexp.MatchString(`^\+([1-9]\d{1,2}|[2-9]\d{2})[ \d\-]{10,16}$`, phone)
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
	match, err := regexp.MatchString(`^((http|https)://)[-a-zA-Z0-9@:%._\\+~#?&//=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%._\\+~#?&//=]*)$
	`, url)
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

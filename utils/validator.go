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

// func ValidateUrl(url string) bool {
// 	match, err := regexp.MatchString(`^(?!mailto:)(?:(?:http|https|ftp)://)(?:\\S+(?::\\S*)?@)?(?:(?:(?:[1-9]\\d?|1\\d\\d|2[01]\\d|22[0-3])(?:\\.(?:1?\\d{1,2}|2[0-4]\\d|25[0-5])){2}(?:\\.(?:[0-9]\\d?|1\\d\\d|2[0-4]\\d|25[0-4]))|(?:(?:[a-z\\u00a1-\\uffff0-9]+-?)*[a-z\\u00a1-\\uffff0-9]+)(?:\\.(?:[a-z\\u00a1-\\uffff0-9]+-?)*[a-z\\u00a1-\\uffff0-9]+)*(?:\\.(?:[a-z\\u00a1-\\uffff]{2,})))|localhost)(?::\\d{2,5})?(?:(/|\\?|#)[^\\s]*)?$
// 	`, url)
// 	if err != nil {
// 		fmt.Errorf("error validating phone number: %w", err)
// 		return false
// 	}
// 	if !match {
// 		fmt.Errorf("invalid phone number format (must start with + and valid international code)")
// 		return false
// 	}
// 	return true
// }

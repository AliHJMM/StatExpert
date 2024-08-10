package functions

import "regexp"
// isValidNumber checks if a string contains only digits.
func isValidNumber(s string) bool {
	match, _ := regexp.MatchString(`^\d+$`, s) // Use regex to check if the string is composed of digits only
	return match
}

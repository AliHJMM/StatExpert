package functions

import "regexp"

func isValidNumber(s string) bool {
	match, _ := regexp.MatchString(`^\d+$`, s)
	return match
}

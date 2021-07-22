package utils

import "regexp"

// IsValidMD5 Returns if a string is a valid MD5 hash
func IsValidMD5(s string) bool {
	matched, err := regexp.Match("^[0-9a-fA-F]{32}$", []byte(s))

	if err != nil {
		return false
	}

	return matched
}

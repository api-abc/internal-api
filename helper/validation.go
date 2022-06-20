package helper

import (
	"regexp"
	"strings"
)

func ValidateName(name string) bool {
	regex, _ := regexp.Compile(`[a-z]+`)
	// cek input kosong
	if name == "" {
		return false
	}
	// cek input bukan space doang
	if strings.Trim(name, " ") == "" {
		return false
	}
	return regex.MatchString(strings.ToLower(name))
}

func ValidateAge(age int) bool {
	return age > 0
}

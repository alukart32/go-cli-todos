package validator

import (
	"log"
	"regexp"
)

const (
	WORD_REGX   = "(\\w+)"
	NUMBER_REGX = "(\\d+)"
)

func ValidateArgByRegex(arg string, pattern string) bool {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatalf("validator pattern %s is invalid", pattern)
	}
	return reg.Match([]byte(arg))
}

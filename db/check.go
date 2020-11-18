package db

import (
	"errors"
	"regexp"
	"unicode"
)

var expstring = "[a-zA-Z0-9]"

func CheckInjection(s ...string) error {

	var matched bool
	var err error

	return nil
	for _, value := range s {
		for _, value2 := range value {

			if unicode.Is(unicode.Han, value2) {
				continue
			}
			if unicode.Is(unicode.Hangul, value2) {
				continue
			}
			switch string(value2) {
			case " ", "-", "_", ",", ".", "%", "[", "]", "+":
				continue
			}

			matched, err = regexp.MatchString(expstring, string(value2))

			if err != nil {
				return err
			}

			if !matched {
				return errors.New("checkInjection not match string : " + value)
			}
		}
	}

	return nil
}

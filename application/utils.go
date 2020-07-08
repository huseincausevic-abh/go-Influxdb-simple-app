package main

import (
	"errors"
	"regexp"
)

func checkEmptyBody(body string) error {
	matched, err := regexp.MatchString(`^\s*$`, body)
	if err != nil {
		panic(err)
	}
	if matched == true {
		return errors.New("Request body is empty")
	}
	return nil
}

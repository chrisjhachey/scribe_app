package scribe

import (
	"errors"
	"regexp"
	"strings"
)

const MAX_NAME_LENGTH = 100

func GenerateURIFromName(name string) (string, error) {
	re := regexp.MustCompile("^[a-zA-Z0-9_ ]*$")

	if !re.MatchString(name) {
		return "", errors.New("name is an invalid format")
	}

	trim := strings.TrimSpace(name)

	if len(trim) < 1 {
		return "", errors.New("no name provided")
	}

	formattedName := ""
	strArray := strings.Fields(strings.ToLower(trim))

	for _, s := range strArray {
		if len(formattedName) < 1 {
			formattedName = s
		} else {
			formattedName = formattedName + "_" + s
		}
	}

	if len(formattedName) > MAX_NAME_LENGTH {
		return "", errors.New("name must be less than 100 characters")
	}

	return formattedName, nil
}

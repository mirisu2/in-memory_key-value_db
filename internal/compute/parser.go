package compute

import (
	"errors"
	"regexp"
	"strings"
)

var (
	validCommands = map[string]bool{
		"SET":    true,
		"GET":    true,
		"DELETE": true,
	}
	allowedChars = `^[a-zA-Z0-9_*/.-]+$`
)

func Parse(request string) (command string, args []string, err error) {
	parts := strings.Fields(request)
	if len(parts) < 2 {
		return "", nil, errors.New("not enough arguments")
	}

	command = parts[0]
	if _, ok := validCommands[command]; !ok {
		return "", nil, errors.New("invalid command")
	}

	argPattern := regexp.MustCompile(allowedChars)
	for _, arg := range parts[1:] {
		if !argPattern.MatchString(arg) {
			return "", nil, errors.New("invalid characters in arguments")
		}
		args = append(args, arg)
	}

	return command, args, nil
}

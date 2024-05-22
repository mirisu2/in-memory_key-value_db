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

	ErrNotEnoughArguments = errors.New("not enough arguments")
	ErrInvalidCommand     = errors.New("invalid command")
	ErrInvalidCharacters  = errors.New("invalid characters")
)

func Parse(request string) (command string, args []string, err error) {
	parts := strings.Fields(request)
	if len(parts) < 2 {
		return "", nil, ErrNotEnoughArguments
	}

	command = parts[0]
	if _, ok := validCommands[command]; !ok {
		return "", nil, ErrInvalidCommand
	}

	argPattern := regexp.MustCompile(allowedChars)
	for _, arg := range parts[1:] {
		if !argPattern.MatchString(arg) {
			return "", nil, ErrInvalidCharacters
		}
		args = append(args, arg)
	}

	return command, args, nil
}

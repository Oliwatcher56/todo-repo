package commands

import (
	"errors"
	"strings"
)

func ExtractQuoted(input string) (string, error) {

	if strings.Count(input, "\"") != 2 {
		return "", errors.New("Аргумент должен быть указан в кавычках: \"task name\"")
	}

	name := input[strings.Index(input, "\"")+1 : strings.LastIndex(input, "\"")]

	return name, nil

}

func ParseAddCommand(input string) (name, priority string, err error) {

	name, err = ExtractQuoted(input)

	if err != nil {
		return "", "", err
	}

	endQuote := strings.LastIndex(input, "\"")
	rightSide := strings.TrimSpace(input[endQuote+1:])

	parts := strings.Fields(rightSide)
	if len(parts) != 1 {
		return "", "", errors.New("Неверный формат ввода, попробуйте: add \"task name\" low|high")
	}

	priority = parts[0]

	return name, priority, nil
}

func ParseChangeCommand(input string) (name, priority, status string, err error) {

	name, err = ExtractQuoted(input)

	if err != nil {
		return "", "", "", err
	}

	endQuote := strings.LastIndex(input, "\"")
	rightSide := strings.TrimSpace(input[endQuote+1:])

	parts := strings.Fields(rightSide)
	if len(parts) != 2 {
		return "", "", "", errors.New("Неверный формат ввода, попробуйте: change \"task name\" low|high todo|progress|done")
	}

	priority = parts[0]
	status = parts[1]

	return name, priority, status, nil
}

func ParseRemoveCommand(input string) (name string, err error) {

	name, err = ExtractQuoted(input)

	if err != nil {
		return "", err
	}

	endQuote := strings.LastIndex(input, "\"")
	rightSide := strings.TrimSpace(input[endQuote+1:])

	if rightSide != "" {
		return "", errors.New("Неверный формат ввода, попробуйте: remove \"task name\"")
	}

	return name, nil
}

func ParseQueryCommand(input string) (query string, err error) {

	query, err = ExtractQuoted(input)

	if err != nil {
		return "", errors.New("Неверный формат ввода, попробуйте: search \"keyword\"")
	}

	endQuote := strings.LastIndex(input, "\"")
	rightSide := strings.TrimSpace(input[endQuote+1:])

	if rightSide != "" {
		return "", errors.New("Неверный формат ввода, попробуйте: search \"keyword\"")
	}

	return query, nil
}

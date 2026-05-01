package commands

import (
	"errors"
	"strings"
)

const (
	EmpytArg              = "Поле аргумента не может быть пустым!"
	InvalidQuoted         = "Аргумент должен быть указан в кавычках: \"argument\""
	InvalidAddCommand     = "Неверный формат ввода, попробуйте: add \"task name\" low|high"
	InvalidChangeCommand  = "Неверный формат ввода, попробуйте: change \"task name\" low|high todo|progress|done"
	InvalidRemoveCommand  = "Неверный формат ввода, попробуйте: remove \"task name\""
	InvalidQueryCommand   = "Неверный формат ввода, попробуйте: search \"keyword\""
	InvalidResetCommand   = "Неверный формат ввода, попробуйте: reset"
	InvalidStorageCommand = "Неверный формат ввода, попробуйте: save|load \"filename.gob\n"
	InvalidFilterCommand  = "Неверный формат ввода, попробуйте: filter \"by low|high|todo|progress|done status|priority\""
)

func ExtractQuoted(input string) (string, error) {

	if strings.Count(input, "\"") != 2 {
		return "", errors.New(InvalidQuoted)
	}

	arg := input[strings.Index(input, "\"")+1 : strings.LastIndex(input, "\"")]

	if arg == "" {
		return "", errors.New(EmpytArg)
	}

	return arg, nil

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
		return "", "", errors.New(InvalidAddCommand)
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
		return "", "", "", errors.New(InvalidChangeCommand)
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
		return "", errors.New(InvalidRemoveCommand)
	}

	return name, nil
}

func ParseQueryCommand(input string) (query string, err error) {

	query, err = ExtractQuoted(input)

	if err != nil {
		return "", errors.New(InvalidQueryCommand)
	}

	endQuote := strings.LastIndex(input, "\"")
	rightSide := strings.TrimSpace(input[endQuote+1:])

	if rightSide != "" {
		return "", errors.New(InvalidQueryCommand)
	}

	return query, nil
}

func ParseResetCommand(input string) error {

	parts := strings.Fields(input)

	if len(parts) != 1 {
		return errors.New(InvalidResetCommand)
	}

	return nil

}

func ParseStorageCommand(input string) (filename string, err error) {

	filename, err = ExtractQuoted(input)

	if err != nil {
		return "", errors.New(InvalidStorageCommand)
	}

	parts := strings.Fields(input)

	if len(parts) > 2 {
		return "", errors.New(InvalidStorageCommand)
	}

	return filename, nil

}

func ParseFilterCommand(input string) (criterion string, mode string, err error) {

	userInput, err := ExtractQuoted(input)

	if err != nil {
		return "", "", errors.New(InvalidFilterCommand)
	}

	parts := strings.Fields(userInput)

	if len(parts) != 3 {
		return "", "", errors.New(InvalidFilterCommand)
	}

	if parts[0] != "by" {
		return "", "", errors.New(InvalidFilterCommand)
	}

	if parts[1] != "todo" && parts[1] != "progress" && parts[1] != "done" && parts[1] != "low" && parts[1] != "high" {
		return "", "", errors.New(InvalidFilterCommand)
	}

	criterion = parts[1]

	if parts[2] != "status" && parts[2] != "priority" {
		return "", "", errors.New(InvalidFilterCommand)
	}

	mode = parts[2]

	endQuote := strings.LastIndex(input, "\"")
	rightSide := strings.TrimSpace(input[endQuote+1:])

	if rightSide != "" {
		return "", "", errors.New(InvalidFilterCommand)
	}

	return criterion, mode, nil

}

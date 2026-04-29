package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Name     string
	Status   string
	Priority string
}

type Commands struct {
	Data  map[string]string
	Order []string
}

var taskList = make(map[string]*Task)

const (
	todo       = "todo"
	inProgress = "progress"
	done       = "done"
	low        = "low"
	high       = "high"
)

const (
	emptyList         = "Список пуст!"
	emptyName         = "Пустое имя задачи!"
	alreadyExists     = "Такая задача уже существует!"
	notExist          = "Такой задачи не существует!"
	protectedByStatus = "Задача в процессе выполнения не может быть удалена!"
	incorrectStatus   = "Некорректный статус задачи"
	incorrectPriority = "Некорректный приоритет задачи"
)

func showList() error {

	if len(taskList) == 0 {
		return errors.New(emptyList)
	}

	for _, task := range taskList {
		fmt.Printf("Задача -> %s | Статус -> %s | Приоритет -> %s\n", task.Name, task.Status, task.Priority)
	}

	return nil
}

func addTask(name, priority string) error {

	name = strings.TrimSpace(name)

	if !isValidName(name) {
		return errors.New(emptyName)
	}

	priority = strings.ToLower(strings.TrimSpace(priority))

	if !isValidPriority(priority) {
		return errors.New(incorrectPriority)
	}

	if _, ok := taskList[name]; ok {
		return errors.New(alreadyExists)
	}

	taskList[name] = &Task{
		Name:     name,
		Status:   todo,
		Priority: priority,
	}

	return nil
}

func removeTask(name string) error {

	name = strings.TrimSpace(name)

	if !isValidName(name) {
		return errors.New(emptyName)
	}

	task, ok := taskList[name]

	if !ok {
		return errors.New(notExist)
	}

	if task.Status == inProgress {
		return errors.New(protectedByStatus)
	}

	delete(taskList, name)

	return nil
}

func changeTask(name, status, priority string) error {

	name = strings.TrimSpace(name)

	if !isValidName(name) {
		return errors.New(emptyName)
	}

	task, ok := taskList[name]

	if !ok {
		return errors.New(notExist)
	}

	status = strings.ToLower(strings.TrimSpace(status))
	priority = strings.ToLower(strings.TrimSpace(priority))

	if !isValidStatus(status) {
		return errors.New(incorrectStatus)
	}

	if !isValidPriority(priority) {
		return errors.New(incorrectPriority)
	}

	task.Status = status
	task.Priority = priority

	return nil
}

func searchByQuery(query string) {
	flag := false

	for name := range taskList {
		ok := strings.Contains(name, query)

		if ok {
			fmt.Printf("Совпадение с %q, возможно вы ищете : %q\n", query, name)
			flag = true
		}
	}

	if !flag {
		fmt.Println("Ничего не найдено")
	}
}

func isValidStatus(status string) bool {
	return status == todo || status == inProgress || status == done
}

func isValidPriority(priority string) bool {
	return priority == low || priority == high
}

func isValidName(name string) bool {
	return name != ""
}

func showCommands(data map[string]string, order []string) {
	for _, name := range order {
		fmt.Printf("Команда -> %s -> %s\n", name, data[name])
	}
}

func main() {

	cmdOrder := []string{"help", "list", "add", "change", "remove", "search", "exit"}

	cmdData := map[string]string{
		"help":   "Вывести список всех команд",
		"list":   "Вывести список всех задач",
		"add":    "Добавить задачу",
		"change": "Изменить имя/статус/приоритет",
		"remove": "Удалить задачу",
		"search": "Найти задачу",
		"exit":   "Выход",
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("ToDo v0.1")

	for {

		fmt.Print("> ")
		scanner.Scan()

		userInput := scanner.Text()

		parts := strings.Fields(userInput)

		command := parts[0]
		args := parts[1:]

		switch command {
		case "help":
			showCommands(cmdData, cmdOrder)
		case "exit":
			fmt.Println("До скорых встреч!")
			return
		case "list":
			if err := showList(); err != nil {
				fmt.Println(err)
			}
		case "add":
			name := userInput[strings.Index(userInput, "\"")+1 : strings.LastIndex(userInput, "\"")]
			priority := args[len(args)-1]

			if err := addTask(name, priority); err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Задача %q с приоритетом %q была добавлена!\n", name, priority)

			}
		case "remove":
			name := userInput[strings.Index(userInput, "\"")+1 : strings.LastIndex(userInput, "\"")]

			if err := removeTask(name); err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Задача %q была удалена!\n", name)
			}
		case "change":
			name := userInput[strings.Index(userInput, "\"")+1 : strings.LastIndex(userInput, "\"")]
			status := args[len(args)-2]
			priority := args[len(args)-1]

			if err := changeTask(name, status, priority); err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Задача %q была изменена!\n", name)
				fmt.Printf("Актуальный статус %q | актуальный приоритет %q\n", status, priority)
			}
		case "search":
			name := userInput[strings.Index(userInput, "\"")+1 : strings.LastIndex(userInput, "\"")]
			searchByQuery(name)
		}
	}

}

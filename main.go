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

var taskList = make(map[string]*Task)

const (
	todo       = "todo"
	inProgress = "in progress"
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
		fmt.Printf("Задача %s | Статус %s | Приоритет %s\n", task.Name, task.Status, task.Priority)
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

func isValidStatus(status string) bool {
	return status == todo || status == inProgress || status == done
}

func isValidPriority(priority string) bool {
	return priority == low || priority == high
}

func isValidName(name string) bool {
	return name != ""
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	scanner.Scan()

	text := scanner.Text()

	if text == "list" {
		showList()
	}

}

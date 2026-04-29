package tasks

import (
	"errors"
	"fmt"
	"strings"
)

type Task struct {
	Name     string
	Status   string
	Priority string
}

const (
	Todo       = "todo"
	InProgress = "progress"
	Done       = "done"

	Low  = "low"
	High = "high"
)

type Store struct {
	tasks map[string]*Task
}

func NewStore() *Store {
	return &Store{
		tasks: make(map[string]*Task),
	}
}

func (s *Store) List() error {

	if len(s.tasks) == 0 {
		return errors.New(EmptyList)
	}

	for _, task := range s.tasks {
		fmt.Printf("Задача -> %s | Статус -> %s | Приоритет -> %s\n", task.Name, task.Status, task.Priority)
	}

	return nil
}

func (s *Store) Add(name, priority string) error {
	name = strings.TrimSpace(name)

	if !isNotEmpty(name) {
		return errors.New(EmptyName)
	}

	priority = strings.ToLower(strings.TrimSpace(priority))

	if !isValidPriority(priority) {
		return errors.New(IncorrectPriority)
	}

	if _, ok := s.tasks[name]; ok {
		return errors.New(AlreadyExists)
	}

	s.tasks[name] = &Task{
		Name:     name,
		Status:   Todo,
		Priority: priority,
	}

	return nil
}

func (s *Store) Remove(name string) error {
	name = strings.TrimSpace(name)

	if !isNotEmpty(name) {
		return errors.New(EmptyName)
	}

	task, ok := s.tasks[name]

	if !ok {
		return errors.New(NotExist)
	}

	if task.Status == InProgress {
		return errors.New(ProtectedByStatus)
	}

	delete(s.tasks, name)

	return nil
}

func (s *Store) Change(name, priority, status string) error {
	name = strings.TrimSpace(name)

	if !isNotEmpty(name) {
		return errors.New(EmptyName)
	}

	task, ok := s.tasks[name]

	if !ok {
		return errors.New(NotExist)
	}

	status = strings.ToLower(strings.TrimSpace(status))
	priority = strings.ToLower(strings.TrimSpace(priority))

	if !isValidStatus(status) {
		return errors.New(IncorrectStatus)
	}

	if !isValidPriority(priority) {
		return errors.New(IncorrectPriority)
	}

	task.Status = status
	task.Priority = priority

	return nil
}

func (s *Store) Search(query string) error {
	query = strings.TrimSpace(query)

	if !isNotEmpty(query) {
		return errors.New(EmptyQuery)
	}

	normalizedQuery := strings.ToLower(query)

	flag := false

	for name := range s.tasks {
		ok := strings.Contains(strings.ToLower(name), normalizedQuery)

		if ok {
			fmt.Printf("Совпадение с %q, возможно вы ищете : %q\n", query, name)
			flag = true
		}
	}

	if !flag {
		fmt.Println("Ничего не найдено")
		return nil
	}

	return nil
}

// Валидация

func isValidStatus(status string) bool {
	return status == Todo || status == InProgress || status == Done
}

func isValidPriority(priority string) bool {
	return priority == Low || priority == High
}

func isNotEmpty(value string) bool {
	return value != ""
}

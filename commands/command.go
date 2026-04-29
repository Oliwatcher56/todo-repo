package commands

import "fmt"

type Commands struct {
	Data  map[string]string
	Order []string
}

var Default = Commands{
	Order: []string{"help", "list", "add", "change", "remove", "search", "exit"},
	Data: map[string]string{
		"help":   "Вывести список всех команд",
		"list":   "Вывести список всех задач",
		"add":    "Добавить задачу",
		"change": "Изменить приоритет/статус",
		"remove": "Удалить задачу",
		"search": "Найти задачу",
		"exit":   "Выход",
	},
}

func (c Commands) Show() {
	for _, name := range Default.Order {
		fmt.Printf("Команда -> %s -> %s\n", name, Default.Data[name])
	}
}

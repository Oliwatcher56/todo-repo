package commands

import "fmt"

type Commands struct {
	Data  map[string]string
	Order []string
}

var Default = Commands{
	Order: []string{"help", "list", "add", "remove", "change", "search", "filter", "reset", "save", "load", "exit"},
	Data: map[string]string{
		"help":   "Вывести список всех команд",
		"list":   "Вывести список всех задач",
		"add":    "Добавить задачу",
		"remove": "Удалить задачу",
		"change": "Изменить приоритет/статус",
		"search": "Найти задачу",
		"filter": "Отфильтровать задачи по критерию",
		"reset":  "Очистить список задач",
		"save":   "Сохранить задачи в отдельный файл",
		"load":   "Загрузить задачи из стороннего файла",
		"exit":   "Выход",
	},
}

func (c Commands) Show() {
	for _, name := range Default.Order {
		fmt.Printf("Команда -> %s -> %s\n", name, Default.Data[name])
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todo/commands"
	"todo/tasks"
)

func main() {

	store := tasks.NewStore()

	err := tasks.Load(&store.Tasks, tasks.MainStorage)

	if err != nil {
		fmt.Println("Ошибка при попытке загрузить список:", err)
		return
	}

	fmt.Println("Данные успешно загружены из файла!")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("ToDo v0.1")

	for {
		fmt.Print("> ")

		scanner.Scan()

		userInput := scanner.Text()
		parts := strings.Fields(userInput)

		if len(parts) < 1 {
			continue
		}

		command := parts[0]

		switch command {
		case "help":
			commands.Default.Show()

		case "exit":
			fmt.Println("До скорых встреч!")
			return

		case "list":

			if err := store.List(); err != nil {
				fmt.Println(err)
			}

		case "add":
			name, priority, err := commands.ParseAddCommand(userInput)

			if err != nil {
				fmt.Println(err)
				continue
			}

			if err := store.Add(name, priority); err != nil {
				fmt.Println(err)
			} else {
				err := tasks.Save(store.Tasks, tasks.MainStorage)

				if err != nil {
					fmt.Println("Ошибка при записи:", err)
				}

				fmt.Printf("Задача %q с приоритетом %q была добавлена!\n", name, priority)

			}

		case "remove":
			name, err := commands.ParseRemoveCommand(userInput)

			if err != nil {
				fmt.Println(err)
				continue
			}

			if err := store.Remove(name); err != nil {
				fmt.Println(err)
			} else {
				err := tasks.Save(store.Tasks, tasks.MainStorage)

				if err != nil {
					fmt.Println("Ошибка при записи:", err)
				}
				fmt.Printf("Задача %q была удалена!\n", name)
			}

		case "change":
			name, priority, status, err := commands.ParseChangeCommand(userInput)

			if err != nil {
				fmt.Println(err)
				continue
			}

			if err := store.Change(name, priority, status); err != nil {
				fmt.Println(err)
			} else {
				err := tasks.Save(store.Tasks, tasks.MainStorage)

				if err != nil {
					fmt.Println("Ошибка при записи:", err)
				}
				fmt.Printf("Задача %q была изменена!\n", name)
				fmt.Printf("Актуальный статус %q | актуальный приоритет %q\n", status, priority)
			}

		case "search":
			query, err := commands.ParseQueryCommand(userInput)

			if err != nil {
				fmt.Println(err)
				continue
			}

			if err := store.Search(query); err != nil {
				fmt.Println(err)
			}

		case "filter":

			criterion, err := commands.ParseFilterCommand(userInput)

			if err != nil {
				fmt.Println(err)
				continue
			}

			switch criterion {
			case "by status":

			case "by priority":

			}

		case "reset":

			if err := commands.ParseResetCommand(userInput); err != nil {
				fmt.Println(err)
				continue
			}

			if ok := tasks.ResetValidate(); !ok {
				continue
			}

			err := tasks.Reset(store.Tasks)

			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("Список задач успешно очищен!")

		case "save":

			filename, err := commands.ParseStorageCommand(userInput)
			if err != nil {
				fmt.Println(err)
				continue
			}

			if err := tasks.Save(store.Tasks, filename); err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("Текущий список успешно сохранён в файл:", filename)

		case "load":

			filename, err := commands.ParseStorageCommand(userInput)
			if err != nil {
				fmt.Println(err)
				continue
			}

			if err := tasks.Load(&store.Tasks, filename); err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("Текущий список успешно загружен из файла:", filename)

		default:
			fmt.Printf("Данной команды не существует, введите help\n")
		}
	}

}

package tasks

import (
	"bufio"
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Save(store map[string]*Task, filename string) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	encoder := gob.NewEncoder(file)
	encoder.Encode(store)

	file.Close()

	return nil
}

const MainStorage = "tasks.gob"

func Load(store *map[string]*Task, filename string) error {

	file, err := os.Open(filename)

	if err != nil {
		return err
	}

	decoder := gob.NewDecoder(file)
	decoder.Decode(store)

	file.Close()

	return nil
}

func ResetValidate() bool {
	fmt.Println("Подтвердите удаление одной из команд -> y/n?")
	scanner := bufio.NewScanner(os.Stdin)

	for {

		scanner.Scan()
		userInput := scanner.Text()
		parts := strings.Fields(userInput)

		if len(parts) != 1 {
			fmt.Println("Неверный формат ввода")
			continue
		}

		switch parts[0] {
		case "n":
			fmt.Println("Удаление успешно отменено!")
			return false
		case "y":
			return true
		}
	}
}

func Reset(store map[string]*Task) error {

	if len(store) == 0 {
		return errors.New(EmptyList)
	}

	for task := range store {
		delete(store, task)
	}

	err := Save(store, MainStorage)

	if err != nil {
		return err
	}

	return nil

}

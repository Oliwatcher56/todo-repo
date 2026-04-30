package tasks

import (
	"encoding/gob"
	"errors"
	"os"
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

func Reset(store map[string]*Task) error {

	if len(store) == 0 {
		return errors.New(EmptyList)
	}

	for task := range store {
		delete(store, task)
	}

	err := Save(store, "tasks.gob")

	if err != nil {
		return err
	}

	return nil

}

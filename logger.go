package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
)

var tasksFileName = "tasks.json"

func logger[T any](msg T) {
	file, _ := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0o600)
	defer file.Close()
	log.SetOutput(file)
	log.Println(msg)
}

func saveJsonToFs(tasks [][]list.Item) {
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		log.Fatal("something went wrong while creating a json")
	}

	file, _ := os.OpenFile(tasksFileName, os.O_CREATE|os.O_RDWR, 0o600)
	defer file.Close()
	log.SetOutput(file)
	log.SetFlags(0)
	log.Println(string(jsonData))
}

func readJsonFromFs() [][]list.Item {
	data, err := os.ReadFile(tasksFileName)
	if err != nil {
		log.Fatal("couldn't open tasks.json")
	}

	var tasks [][]task

	if err := json.Unmarshal(data, &tasks); err != nil {
		log.Fatal("couldn't unmarshal tasks.json")
	}

	listItems := make([][]list.Item, len(tasks))

	for i, tasksArr := range tasks {
		listItems[i] = make([]list.Item, len(tasksArr))
		for j, task := range tasksArr {
			listItems[i][j] = task
		}
	}

	return listItems
}

package main

import (
	"log"
	"os"
)

func logger [T any](msg T) {
  file, _ := os.OpenFile("deubg.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
  defer file.Close()
  log.SetOutput(file)
  log.Println(msg)
} 

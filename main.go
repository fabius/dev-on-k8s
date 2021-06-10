package main

import (
	"log"
	"os"
	"time"
)

func main() {
	for {
		log.Println("Hello from service " + os.Getenv("NAME"))
		time.Sleep(time.Second)
	}
}

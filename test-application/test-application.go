package main

import (
	"log"
	"time"
)

func main() {
	counter := 1

	for {
		log.Printf("test %d", counter)
		counter++
		time.Sleep(1 * time.Second)
	}
}

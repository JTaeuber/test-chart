package main

import (
	"log/slog"
	"time"
)

func main() {
	var counter int

	for {
		slog.Info("counting up", "count", counter)
		counter++
		time.Sleep(1 * time.Second)
	}
}

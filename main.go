package main

import (
	"bufio"
	"os"

	"github.com/itelman/cinema-ticket-system/internal/handler"
	"github.com/itelman/cinema-ticket-system/internal/repository"
)

func main() {
	cinemaSystem := repository.CinemaSystem()
	handlers := handler.Handlers{System: &cinemaSystem}

	reader := bufio.NewReader(os.Stdin)

	for {
		status := handlers.MainMenu(reader)
		if status == 1 {
			break
		}
	}
}

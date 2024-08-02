package handler

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/itelman/cinema-ticket-system/internal/repository"
)

type Handlers struct {
	System *repository.CinemaTicketSystem
}

func (h *Handlers) ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = log.Writer()
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to clear screen: %v", err)
	}
}

func (h *Handlers) SetDelay() {
	time.Sleep(3 * time.Second)
}

func (h *Handlers) MainMenu(reader *bufio.Reader) int {
	menu := []string{"Здравствуйте, у вас есть следующие доступные функции:\n\n",
		"1. Добавить новый фильм\n",
		"2. Показать все доступные фильмы\n",
		"3. Добавить нового пользователя\n",
		"4. Купить билет\n",
		"5. Отменить покупку билета\n",
		"6. Выход\n\n",
		"Выберите опцию: ",
	}

	for _, lines := range menu {
		fmt.Print(lines)
	}

	option, _ := reader.ReadString('\n')
	option = strings.TrimSpace(option)
	fmt.Println()

	switch option {
	case "1":
		h.AddMovieHandler(reader)
	case "2":
		h.ShowAllMoviesHandler(reader)
	case "3":
		h.AddUserHandler(reader)
	case "4":
		h.BuyTicketHandler(reader)
	case "5":
		h.CancelTicketHandler(reader)
	case "6":
		fmt.Println("Выход из программы.")
		h.SetDelay()
		return 1
	default:
		fmt.Println("Неверная опция. Пожалуйста, выберите правильную опцию.")
	}

	h.SetDelay()
	h.ClearScreen()
	return 0
}

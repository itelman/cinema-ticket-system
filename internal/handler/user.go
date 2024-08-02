package handler

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/itelman/cinema-ticket-system/pkg/response"
)

func (h *Handlers) AddUserHandler(reader *bufio.Reader) {
	fmt.Print("Введите имя пользователя (или введите -1 чтобы вернуться в меню): ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	if name == "-1" {
		return
	}

	id, err := h.System.AddUser(name)
	if err != nil {
		response.RecordExistsError()
		h.SetDelay()
		h.AddUserHandler(reader)
		return
	}

	fmt.Printf("\nПользователь добавлен. ID пользователя: %d\n", id)
}

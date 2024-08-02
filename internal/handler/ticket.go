package handler

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/itelman/cinema-ticket-system/pkg/response"
	"github.com/itelman/cinema-ticket-system/pkg/store"
)

func (h *Handlers) BuyTicketHandler(reader *bufio.Reader) {
	if len(h.System.Users) == 0 || len(h.System.Movies) == 0 {
		fmt.Println("На данный момент невозможно оформить билеты, поскольку нет списка доступных пользователей/фильмов.")
		h.SetDelay()
		return
	}

	fmt.Print("Введите ID пользователя (или введите -1 чтобы вернуться в меню): ")
	userID, _ := reader.ReadString('\n')
	userID = strings.TrimSpace(userID)

	if userID == "-1" {
		return
	}

	user_id, err := strconv.Atoi(userID)
	if err != nil {
		response.BadRequestError()
		h.SetDelay()
		h.BuyTicketHandler(reader)
		return
	}

	fmt.Print("Введите ID фильма (или введите -1 чтобы вернуться в меню): ")
	movieID, _ := reader.ReadString('\n')
	movieID = strings.TrimSpace(movieID)

	if movieID == "-1" {
		return
	}

	movie_id, err := strconv.Atoi(movieID)
	if err != nil {
		response.BadRequestError()
		h.SetDelay()
		h.BuyTicketHandler(reader)
		return
	}

	id, err := h.System.BuyTicket(user_id, movie_id)
	if err != nil {
		if err == store.ErrNotFound {
			response.NotFoundError()
		} else if err == store.ErrRecordExists {
			response.RecordExistsError()
		}
		h.SetDelay()
		h.BuyTicketHandler(reader)
		return
	}

	fmt.Printf("\nБилет успешно оформлен. ID билета: %d\n", id)
}

func (h *Handlers) CancelTicketHandler(reader *bufio.Reader) {
	if len(h.System.Tickets) == 0 {
		fmt.Println("На данный момент не оформлены никакие билеты.")
		h.SetDelay()
		return
	}

	fmt.Print("Введите ID билета (или введите -1 чтобы вернуться в меню): ")
	ticketID, _ := reader.ReadString('\n')
	ticketID = strings.TrimSpace(ticketID)

	if ticketID == "-1" {
		return
	}

	ticket_id, err := strconv.Atoi(ticketID)
	if err != nil {
		response.BadRequestError()
		h.SetDelay()
		h.CancelTicketHandler(reader)
		return
	}

	exists := h.System.CancelTicket(ticket_id)
	if !exists {
		response.NotFoundError()
		h.SetDelay()
		h.CancelTicketHandler(reader)
		return
	}

	fmt.Println("\nПокупка билета отменена.")
}

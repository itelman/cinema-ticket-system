package repository

import (
	"strings"

	"github.com/itelman/cinema-ticket-system/internal/repository/entities"
	"github.com/itelman/cinema-ticket-system/pkg/store"
)

type CinemaTicketSystem struct {
	Movies  map[int]*entities.Movie
	Users   map[int]*entities.User
	Tickets map[int]*entities.Ticket
}

func CinemaSystem() CinemaTicketSystem { // constructor was named "CinemaSystem" instead of "CinemaTicketSystem" due to Golang's redeclaring conflict
	return CinemaTicketSystem{
		make(map[int]*entities.Movie),
		make(map[int]*entities.User),
		make(map[int]*entities.Ticket),
	}
}

func (c *CinemaTicketSystem) AddMovie(movieName string) (int, error) {
	movieName = strings.Title(strings.ToLower(movieName))

	for _, movie := range c.Movies {
		if movie.Title == movieName {
			return -1, store.ErrRecordExists
		}
	}

	id := len(c.Movies) + 1
	newMovie := entities.NewMovie(id, movieName)
	c.Movies[id] = &newMovie

	return id, nil
}

func (c *CinemaTicketSystem) AddUser(userName string) (int, error) {
	userName = strings.ToLower(userName)

	for _, user := range c.Users {
		if user.Name == userName {
			return -1, store.ErrRecordExists
		}
	}

	id := len(c.Users) + 1
	newUser := entities.NewUser(id, userName)
	c.Users[id] = &newUser

	return id, nil
}

func (c *CinemaTicketSystem) ShowAllMovies() map[int]*entities.Movie {
	return c.Movies
}

func (c *CinemaTicketSystem) BuyTicket(userId int, movieId int) (int, error) {
	if _, ok := c.Users[userId]; !ok {
		return -1, store.ErrNotFound
	}

	if _, ok := c.Movies[movieId]; !ok {
		return -1, store.ErrNotFound
	}

	for _, ticket := range c.Tickets {
		if ticket.UserID == userId && ticket.MovieID == movieId {
			return -1, store.ErrRecordExists
		}
	}

	id := len(c.Tickets) + 1
	newTicket := entities.NewTicket(id, userId, movieId)
	c.Tickets[id] = &newTicket

	return id, nil
}

func (c *CinemaTicketSystem) CancelTicket(ticketId int) bool {
	if _, ok := c.Tickets[ticketId]; !ok {
		return false
	}

	delete(c.Tickets, ticketId)
	return true
}

package handler

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/itelman/cinema-ticket-system/pkg/response"
)

func (h *Handlers) AddMovieHandler(reader *bufio.Reader) {
	fmt.Print("Введите название фильма (или введите -1 чтобы вернуться в меню): ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	if title == "-1" {
		return
	}

	id, err := h.System.AddMovie(title)
	if err != nil {
		response.RecordExistsError()
		h.SetDelay()
		h.AddMovieHandler(reader)
		return
	}

	fmt.Printf("\nФильм добавлен. ID фильма: %d\n", id)
}

func (h *Handlers) ShowAllMoviesHandler(reader *bufio.Reader) {
	movies := h.System.ShowAllMovies()

	if len(movies) == 0 {
		fmt.Println("На данный момент нет доступных фильмов.")
		return
	}

	fmt.Println("Доступные фильмы:\n")
	for _, movie := range movies {
		fmt.Printf("%d. %s\n", movie.ID, movie.Title)
	}
}

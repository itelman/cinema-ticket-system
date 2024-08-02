package entities

type Movie struct {
	ID    int
	Title string
}

func NewMovie(id int, title string) Movie {
	return Movie{id, title}
}

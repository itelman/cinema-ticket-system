package entities

type Ticket struct {
	ID      int
	UserID  int
	MovieID int
}

func NewTicket(id, userId, movieId int) Ticket {
	return Ticket{id, userId, movieId}
}

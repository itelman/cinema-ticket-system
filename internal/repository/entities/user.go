package entities

type User struct {
	ID   int
	Name string
}

func NewUser(id int, name string) User {
	return User{id, name}
}

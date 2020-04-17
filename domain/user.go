package domain

type Users []User

type User struct {
	ID       int
	Name     string
	Password string
	Email    string
}

type LoginUser struct {
	ID       int
	Password string
}

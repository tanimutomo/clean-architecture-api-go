package domain

type Articles []Article

type Article struct {
	ID      int
	Title   string
	Content string
	UserID  int
}

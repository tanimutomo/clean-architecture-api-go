package domain

type Tags []Tag

type Tag struct {
	ID        int
	Name      string
	ArticleID int
}

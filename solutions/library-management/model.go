package main

type Category int

const (
	Horror Category = iota
	Romance
	Fiction
	SciFi
)

type BookStatus int

const (
	Available BookStatus = iota
	Taken
)

type SearchType int

const (
	ByAuthor SearchType = iota
	ByCategory
)

type Book struct {
	Id            string
	ISBN          string
	Title         string
	AuthorId      string
	PublicationId string
	Year          int
	Categories    []Category
	Status        BookStatus
}

type Author struct {
	Id   string
	Name string
}

type Publication struct {
	Id   string
	Name string
}

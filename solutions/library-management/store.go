package main

type Store interface {
	GetBooks() []Book
	GetAuthors() []Author
	UpdateBookStatus(string, BookStatus)
}

type StoreImpl struct {
	books        []Book
	authors      []Author
	publications []Publication
}

func (s *StoreImpl) GetBooks() []Book {
	return s.books
}

func (s *StoreImpl) GetAuthors() []Author {
	return s.authors
}

func (s *StoreImpl) UpdateBookStatus(bookId string, status BookStatus) {
	for i := 0; i < len(s.books); i += 1 {
		if s.books[i].Id == bookId {
			s.books[i].Status = status
		}
	}
}

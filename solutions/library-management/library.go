package main

type Library interface {
	Search(params map[SearchType]interface{}) []Book
	Reserve(bookId string)
	Return(bookId string, userId string) int
	History(userId string)
}

type LibraryImpl struct {
	searchService SearchService
}

func (l LibraryImpl) Search(params map[SearchType]interface{}) []Book {
	return l.searchService.Search(params)
}

func (l LibraryImpl) Reserve(bookId string) {
	//TODO implement me
	panic("implement me")
}

func (l LibraryImpl) Return(bookId string, userId string) int {
	//TODO implement me
	panic("implement me")
}

func (l LibraryImpl) History(userId string) {
	//TODO implement me
	panic("implement me")
}

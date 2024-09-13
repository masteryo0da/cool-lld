package main

import (
	"fmt"
	"slices"
)

type FilterAction interface {
	Filter([]Book) []Book
}

type CategoryFilter struct {
	category Category
}

func (c CategoryFilter) Filter(books []Book) []Book {
	result := make([]Book, 0)
	for _, book := range books {
		if slices.Contains(book.Categories, c.category) {
			result = append(result, book)
		}
	}
	return result
}

type AuthorFilter struct {
	store Store
	term  string
}

func (a AuthorFilter) Filter(books []Book) []Book {
	result := make([]Book, 0)
	for _, author := range a.store.GetAuthors() {
		if author.Name == a.term {
			for _, book := range books {
				if book.AuthorId == author.Id {
					result = append(result, book)
				}
			}
		}
	}
	return result
}

type FilterActionFactory struct {
	store Store
}

func (f FilterActionFactory) GetFilterAction(searchType SearchType, metaData interface{}) (FilterAction, error) {

	switch searchType {
	case ByAuthor:
		authorTerm, ok := metaData.(string)
		if ok {
			return AuthorFilter{
				store: f.store,
				term:  authorTerm,
			}, nil
		}
	case ByCategory:
		category, ok := metaData.(Category)
		if !ok {
			return CategoryFilter{category: category}, nil
		}
	}
	return nil, fmt.Errorf("error creating filter action")
}

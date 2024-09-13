package main

func main() {
	authors := []Author{
		{"auth-1", "J.R.R Tolkien"},
	}
	publications := []Publication{
		{"pub-1", "HarperCollins"},
	}
	books := []Book{
		{"book-1", "isbn-1", "Lord of the rings - 1", "auth-1",
			"pub-1", 1954, []Category{Fiction}, Available},
		{"book-2", "isbn-2", "Lord of the rings - 2", "auth-1",
			"pub-1", 1954, []Category{Fiction}, Available},
	}
	store := &StoreImpl{books: books, authors: authors, publications: publications}
	factory := FilterActionFactory{store: store}
	searchService := SearchServiceImpl{
		store:         store,
		filterFactory: factory,
	}
	library := LibraryImpl{searchService: searchService}
	result := library.Search(map[SearchType]interface{}{
		ByAuthor: "J.R.R Tolkien",
	})
	print(result)
}

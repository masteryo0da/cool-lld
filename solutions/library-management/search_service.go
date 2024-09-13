package main

type SearchService interface {
	Search(params map[SearchType]interface{}) []Book
}

type SearchServiceImpl struct {
	store         Store
	filterFactory FilterActionFactory
}

func (s SearchServiceImpl) Search(params map[SearchType]interface{}) []Book {
	result := s.store.GetBooks()
	for param, value := range params {
		filterAction, err := s.filterFactory.GetFilterAction(param, value)
		if err != nil {
			return result
		}
		result = filterAction.Filter(result)
	}
	return result
}

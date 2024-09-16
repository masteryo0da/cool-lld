package main

type SearchService struct {
	store Store
}

func (s SearchService) Search(query SearchQuery) []Show {
	shows := make([]Show, 0)
	for _, show := range s.store.shows {
		movie := s.store.GetMovieById(show.MovieId)
		cinemaId := s.store.GetCinemaById(show.CinemaId)
		if movie != nil {
			isPossible := false
			if query.Title != nil && *query.Title == movie.Name {
				isPossible = true
			} else if query.Title != nil && *query.Genre == movie.Genre {
				isPossible = true
			} else if query.Language != nil && *query.Language == movie.Language {
				isPossible = true
			}
			if isPossible && cinemaId.CityId == query.City {
				shows = append(shows, show)
			}
		}
	}
	return shows
}

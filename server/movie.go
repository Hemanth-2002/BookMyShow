package server

import (
	pb "bms/bmsproto"
	imdb "bms/imdbRating"
	"bms/model"
	"context"
	"log"
)

// function to add new movie on server
func (s *BmsServer) AddMovie(ctx context.Context, in *pb.NewMovie) (*pb.Movie, error) {
	log.Printf("creating new movie called")
	newMovie := model.Movie{
		MovieName:   in.GetMovieName(),
		Director:    in.GetDirector(),
		Description: in.GetDescription(),
		Language:    in.GetLanguage(),
		Genre:       in.GetGenre(),
		ReleaseDate: in.GetReleaseDate(),
		Status:      in.GetStatus(),
	}
	//  if no rating is given then rating is fetched using public api else the given rating is taken
	if in.GetRating() == 0 {
		newMovie.Rating = imdb.ImdbRating(in.GetMovieName())
		s.Db.AddMovie(newMovie)
		return &pb.Movie{
			MovieName:   in.GetMovieName(),
			Director:    in.GetDirector(),
			Description: in.GetDescription(),
			Rating:      uint64(imdb.ImdbRating(in.GetMovieName())),
			Language:    in.GetLanguage(),
			Genre:       in.GetGenre(),
			ReleaseDate: in.GetReleaseDate(),
			Status:      in.GetStatus(),
			Id:          uint64(newMovie.ID)}, nil
	} else {
		newMovie.Rating = int(in.GetRating())
		s.Db.AddMovie(newMovie)
		return &pb.Movie{
			MovieName:   in.GetMovieName(),
			Director:    in.GetDirector(),
			Description: in.GetDescription(),
			Rating:      in.Rating,
			Language:    in.GetLanguage(),
			Genre:       in.GetGenre(),
			ReleaseDate: in.GetReleaseDate(),
			Status:      in.GetStatus(),
			Id:          uint64(newMovie.ID)}, nil
	}
}

// function to get all movies on server
func (s *BmsServer) GetMovies(ctx context.Context, in *pb.EmptyMovie) (*pb.Movies, error) {
	log.Printf("Getting employees called")
	AllMovies := []*pb.Movie{}
	Movies, err := s.Db.GetMovies()
	CheckCall(err)
	for _, movie := range Movies {
		AllMovies = append(AllMovies, &pb.Movie{
			MovieName:   movie.MovieName,
			Director:    movie.Director,
			Description: movie.Description,
			Rating:      uint64(movie.Rating),
			Language:    movie.Language,
			Genre:       movie.Genre,
			ReleaseDate: movie.ReleaseDate,
			Status:      movie.Status,
		})
	}
	return &pb.Movies{Movies: AllMovies}, nil
}

// function to get movie by preference (rating,language,genre)
func (s *BmsServer) GetMovieByPreference(ctx context.Context, in *pb.Movie) (*pb.Movies, error) {
	log.Printf("Getting movie by preference called")
	AllMovies := []*pb.Movie{}
	language := in.GetLanguage()
	rating := in.GetRating()
	genre := in.GetGenre()
	movie := model.Movie{Language: language, Rating: int(rating), Genre: genre}
	Movies, err := s.Db.GetMovie(movie)
	CheckCall(err)
	for _, movie := range Movies {
		AllMovies = append(AllMovies, &pb.Movie{
			MovieName:   movie.MovieName,
			Director:    movie.Director,
			Description: movie.Description,
			Rating:      uint64(movie.Rating),
			Language:    movie.Language,
			Genre:       movie.Genre,
			ReleaseDate: movie.ReleaseDate,
			Status:      movie.Status,
		})
	}
	return &pb.Movies{Movies: AllMovies}, nil
}

// function to update movie status on server
func (s *BmsServer) UpdateMovieStatus(ctx context.Context, in *pb.Movie) (*pb.Movie, error) {
	log.Printf("update movie status called")
	updatedStatus := model.Movie{
		Status: in.GetStatus(),
	}
	updatedStatus.ID = uint(in.Id)
	err := s.Db.UpdateMovie(updatedStatus)
	CheckCall(err)
	return &pb.Movie{Status: in.GetStatus()}, nil
}

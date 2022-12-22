package server

import (
	pb "bms/bmsproto"
	"bms/mocks"
	model "bms/model"
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

var MovieUpdate = model.Movie{
	Status: "Active",
}

var MockMovieUpdate = pb.Movie{
	Status: "Active",
}

var NewMovie = pb.NewMovie{
	MovieName:   "Avatar",
	Director:    "Cameron",
	Description: "Movie description ....",
	Rating:      9,
	Language:    "English",
	Genre:       "suspense",
	ReleaseDate: "1-3-2022",
	Status:      "Active",
}

var MockMovie = model.Movie{
	MovieName:   "Avatar",
	Director:    "Cameron",
	Description: "Movie description ....",
	Rating:      9,
	Language:    "English",
	Genre:       "suspense",
	ReleaseDate: "1-3-2022",
	Status:      "Active",
}

var MockPreferredMovie = model.Movie{
	Genre:  "Action",
	Rating: 9,
}

var PreferredMovie = pb.Movie{
	Genre:  "Action",
	Rating: 9,
}

func TestAddMovie(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBase(controller)
	testMovie := BmsServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().AddMovie(MockMovie).Return(nil)
	got, err := testMovie.AddMovie(ctx, &NewMovie)
	model.CheckError(err)
	expected := &pb.Movie{
		MovieName:   "Avatar",
		Director:    "Cameron",
		Description: "Movie description ....",
		Rating:      9,
		Language:    "English",
		Genre:       "suspense",
		ReleaseDate: "1-3-2022",
		Status:      "Active",
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

func TestGetMovies(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBase(controller)
	testMovie := BmsServer{Db: mockDb}
	ctx := context.Background()

	mockDb.EXPECT().GetMovies().Return([]model.Movie{MockMovie}, nil)
	movies, err := testMovie.GetMovies(ctx, &pb.EmptyMovie{})
	var got = movies.Movies
	model.CheckError(err)
	expected := []*pb.Movie{}
	expected = append(expected, &pb.Movie{
		MovieName:   "Avatar",
		Director:    "Cameron",
		Description: "Movie description ....",
		Rating:      9,
		Language:    "English",
		Genre:       "suspense",
		ReleaseDate: "1-3-2022",
		Status:      "Active",
	})
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

func TestGetMovieByPreference(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBase(controller)
	testMovie := BmsServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().GetMovie(MockPreferredMovie).Return([]model.Movie{MockMovie}, nil)
	movies, err := testMovie.GetMovieByPreference(ctx, &PreferredMovie)
	var got = movies.Movies
	model.CheckError(err)
	expected := []*pb.Movie{}
	expected = append(expected, &pb.Movie{
		MovieName:   "Avatar",
		Director:    "Cameron",
		Description: "Movie description ....",
		Rating:      9,
		Language:    "English",
		Genre:       "suspense",
		ReleaseDate: "1-3-2022",
		Status:      "Active",
	})

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

func TestUpdateMovieStatus(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBase(controller)
	testMovie := BmsServer{Db: mockDb}
	ctx := context.Background()
	MovieUpdate.ID = 1
	mockDb.EXPECT().UpdateMovie(MovieUpdate).Return(nil)
	MockMovieUpdate.Id = 1
	got, err := testMovie.UpdateMovieStatus(ctx, &MockMovieUpdate)

	model.CheckError(err)
	expected := &pb.Movie{
		Status: "Active",
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

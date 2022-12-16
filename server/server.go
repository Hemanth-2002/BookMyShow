package main

import (
	pb "bms/bmsproto"
	imdb "bms/imdbRating"
	model "bms/model"
	"context"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

const (
	port = ":50051" // choosing port number
)

type BmsServer struct {
	pb.UnimplementedBmsDatabaseCrudServer
	Db *gorm.DB
}

// function to add new user on server
func (s *BmsServer) CreateUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("creating new user called")
	newUser := model.User{
		UserName:    in.GetUserName(),
		Password:    in.GetPassword(),
		Email:       in.GetEmail(),
		PhoneNumber: in.GetPhoneNumber(),
	}
	s.Db.Save(&newUser)
	return &pb.User{UserName: in.GetUserName(), Password: in.GetPassword(), Email: in.GetEmail(), PhoneNumber: in.GetPhoneNumber(), Id: uint64(newUser.ID)}, nil
}

// function to add new movie on server
func (s *BmsServer) AddMovie(ctx context.Context, in *pb.NewMovie) (*pb.Movie, error) {
	log.Printf("creating new movie called")
	newMovie := model.Movie{
		MovieName:   in.GetMovieName(),
		Director:    in.GetDirector(),
		Description: in.GetDescription(),
		Rating:      imdb.ImdbRating(in.GetMovieName()),
		Language:    in.GetLanguage(),
		Genre:       in.GetGenre(),
		ReleaseDate: in.GetReleaseDate(),
		Status:      in.GetStatus(),
	}
	s.Db.Save(&newMovie)
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
}

// function to add new booking on server
func (s *BmsServer) AddBooking(ctx context.Context, in *pb.NewBooking) (*pb.Booking, error) {
	log.Printf("creating new booking called")
	newBooking := model.Booking{
		UserID: int(in.GetUserId()),
		ShowID: int(in.GetShowId()),
		Amount: int(in.GetAmount()),
	}
	s.Db.Save(&newBooking)
	return &pb.Booking{UserId: in.GetUserId(), ShowId: in.GetShowId(), Amount: in.GetAmount(), Id: uint64(newBooking.ID)}, nil
}

// function to add payment for a booking
func (s *BmsServer) AddPayment(ctx context.Context, in *pb.NewPayment) (*pb.Payment, error) {
	log.Printf("creating new booking called")
	NewPayment := model.Payment{
		BookingID:        int(in.GetBookingId()),
		Amount:           int(in.GetAmount()),
		DiscountCouponID: int(in.GetDiscountCouponId()),
		Mode:             in.GetMode(),
		Status:           in.GetStatus(),
	}
	s.Db.Save(&NewPayment)
	return &pb.Payment{Amount: in.GetAmount(), DiscountCouponId: in.GetDiscountCouponId(), Mode: in.GetMode(), Status: in.GetStatus(), Id: uint64(NewPayment.ID)}, nil
}

// function to get all movies on server
func (s *BmsServer) GetMovies(ctx context.Context, in *pb.EmptyMovie) (*pb.Movies, error) {
	log.Printf("Getting employees called")
	Movies := []model.Movie{}
	AllMovies := []*pb.Movie{}
	s.Db.Find(&Movies)
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
func (s *BmsServer) GetMovieByPreference(ctx context.Context, in *pb.MoviePreference) (*pb.Movies, error) {
	log.Printf("Getting movie by preference called")
	Movies := []model.Movie{}
	AllMovies := []*pb.Movie{}
	language := in.GetLanguage()
	rating := in.GetRating()
	genre := in.GetGenre()
	s.Db.Where(&model.Movie{Language: language, Rating: int(rating), Genre: genre}).Find(&Movies)
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

// function to get list of shows by theatre
func (s *BmsServer) GetListOfShowsByTheatre(ctx context.Context, in *pb.Theatre) (*pb.Shows, error) {
	log.Printf("Getting list of shows by theatre called")
	Shows := []model.Show{}
	AllShows := []*pb.Show{}
	TheatreId := in.GetTheatreId()
	s.Db.Where(&model.Show{TheatreID: int(TheatreId)}).Find(&Shows)
	for _, show := range Shows {
		AllShows = append(AllShows, &pb.Show{
			Date:      show.Date,
			StartTime: show.StartTime,
			EndTime:   show.EndTime,
			MovieId:   uint64(show.MovieID),
			TheatreId: uint64(show.TheatreID),
		})
	}
	return &pb.Shows{Shows: AllShows}, nil
}

// function to get list of bookings by user
func (s *BmsServer) GetListOfBookingsByUser(ctx context.Context, in *pb.User) (*pb.Bookings, error) {
	log.Printf("Getting list of bookings by user called")
	Bookings := []model.Booking{}
	AllBookings := []*pb.Booking{}
	UserId := in.GetId()
	s.Db.Where(&model.Booking{UserID: int(UserId)}).Find(&Bookings)
	for _, booking := range Bookings {
		AllBookings = append(AllBookings, &pb.Booking{
			UserId: uint64(booking.UserID),
			ShowId: uint64(booking.ShowID),
			Amount: uint64(booking.Amount),
		})
	}
	return &pb.Bookings{Bookings: AllBookings}, nil
}

// function to update user info on server
func (s *BmsServer) UpdateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	log.Printf("update user info called")
	s.Db.Model(&model.User{}).Where("id=?", in.Id).Updates(model.User{
		Password:    in.GetPassword(),
		Email:       in.GetEmail(),
		PhoneNumber: in.GetPhoneNumber(),
	})
	return &pb.User{UserName: in.GetUserName(), Password: in.GetPassword(), Email: in.GetEmail(), PhoneNumber: in.GetPhoneNumber()}, nil
}

// function to update movie status on server
func (s *BmsServer) UpdateMovieStatus(ctx context.Context, in *pb.Movie) (*pb.Movie, error) {
	log.Printf("update movie status called")
	s.Db.Model(&model.Movie{}).Where("id=?", in.Id).Updates(model.Movie{
		Status: in.GetStatus(),
	})
	return &pb.Movie{Status: in.GetStatus()}, nil
}

func main() {
	model.StartDB()
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err.Error())
	}

	//db connection
	db, err := gorm.Open("postgres", "user=postgres password=MohanNeelima@01 dbname=BookMyShow sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//create new server
	new_server := grpc.NewServer()
	pb.RegisterBmsDatabaseCrudServer(new_server, &BmsServer{
		Db: db,
	})

	log.Printf("Using port no %v", listen.Addr())

	if err := new_server.Serve(listen); err != nil {
		log.Fatal(err.Error())
	}
}

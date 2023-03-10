package main

import (
	pb "bms/bmsproto"
	"bms/utils"
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address = "bms-vog9.onrender.com:50051" // port address for client
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()
	type credentialsl struct {
		credentials.TransportCredentials
		once  sync.Once
		state tls.ConnectionState
	}
	creds := &credentialsl{TransportCredentials: credentials.NewTLS(&tls.Config{})}
	conn, err := grpc.DialContext(ctx, address,
		grpc.WithTransportCredentials(creds),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)))
	fmt.Println("unable to dial")
	utils.CheckError(err)
	defer conn.Close()
	fmt.Println("connected")

	client := pb.NewBmsDatabaseCrudClient(conn)
	//creating new user
	newUser := pb.NewUser{UserName: "PST", Password: "123", Email: "pst@bc.in", PhoneNumber: "9123"}
	new_user, err := client.CreateUser(ctx, &newUser)
	utils.CheckError(err)
	token, err := client.CreateToken(ctx, &newUser)

	utils.CheckError(err)

	log.Printf("User Name: %v, Password: %v, Email: %v, PhoneNumber : %v", new_user.GetUserName(), new_user.GetPassword(), new_user.GetEmail(), new_user.GetPhoneNumber())
	log.Printf("User Token: %v", token)
	// creating a new movie
	new_movie, err := client.AddMovie(ctx, &pb.NewMovie{
		MovieName:   "Avatar",
		Director:    "Cameron",
		Description: "Movie description ....",
		Rating:      9,
		Language:    "English",
		Genre:       "suspense",
		ReleaseDate: "1-3-2022",
		Status:      "Active",
	})
	utils.CheckError(err)

	log.Printf("Movie Name: %v, Director: %v, Description: %v, Rating : %v, Language : %v, Genre : %v, Release date : %v, Status : %v", new_movie.GetMovieName(), new_movie.GetDirector(), new_movie.GetDescription(), new_movie.GetRating(), new_movie.GetLanguage(), new_movie.GetGenre(), new_movie.GetReleaseDate(), new_movie.GetStatus())

	// creating a new booking
	new_booking, err := client.AddBooking(ctx, &pb.NewBooking{
		UserId: 2,
		ShowId: 1,
		Amount: 150,
	})
	utils.CheckError(err)

	log.Printf("User Id: %v, Show Id: %v, Amount: %v", new_booking.GetUserId(), new_booking.GetShowId(), new_booking.GetAmount())

	//adding seat to booking
	new_seat, err := client.AddSeat(ctx, &pb.NewSeat{
		BookingId:  1,
		SeatNumber: 2,
	})
	utils.CheckError(err)

	log.Printf("Booking Id: %v, Seat Number %v", new_seat.GetBookingId(), new_seat.GetSeatNumber())

	// creating a new payment
	new_payment, err := client.AddPayment(ctx, &pb.NewPayment{
		Amount:    150,
		Mode:      "credit-card",
		Status:    true,
		BookingId: 2,
		UserId:    1,
	})
	utils.CheckError(err)

	log.Printf("Amount: %v, Mode: %v, Status: %v", new_payment.GetAmount(), new_payment.GetMode(), new_payment.GetStatus())

	// Get all movies
	AllMovies, err := client.GetMovies(ctx, &pb.EmptyMovie{})
	utils.CheckError(err)

	for _, movie := range AllMovies.Movies {
		fmt.Println(movie.GetMovieName(), movie.GetDirector(), movie.GetDescription(), movie.GetRating(), movie.GetLanguage(), movie.GetGenre(), movie.GetReleaseDate(), movie.GetStatus())
	}

	// get movie by preference
	MovieByPreference, err := client.GetMovieByPreference(ctx, &pb.Movie{
		Genre:  "Action",
		Rating: 8,
	})
	utils.CheckError(err)

	for _, movie := range MovieByPreference.Movies {
		fmt.Println(movie.GetMovieName(), movie.GetDirector(), movie.GetDescription(), movie.GetRating(), movie.GetLanguage(), movie.GetGenre(), movie.GetReleaseDate(), movie.GetStatus())
	}

	// get list of shows by theatre
	ShowsList, err := client.GetListOfShowsByTheatre(ctx, &pb.Theatre{
		TheatreId: 1,
	})
	utils.CheckError(err)

	for _, show := range ShowsList.Shows {
		fmt.Println(show.GetDate(), show.GetStartTime(), show.GetEndTime(), show.GetMovieId(), show.GetTheatreId())
	}

	// get list of bookings by user
	BookingsList, err := client.GetListOfBookingsByUser(ctx, &pb.User{
		Id: 1,
	})
	utils.CheckError(err)

	for _, booking := range BookingsList.Bookings {
		fmt.Println(booking.GetUserId(), booking.GetShowId(), booking.GetAmount())
	}

	// updating user info
	updated_user, err := client.UpdateUser(ctx, &pb.User{Password: "newPwd", PhoneNumber: "456", UserName: "GST", Id: 3})
	utils.CheckError(err)
	fmt.Println(updated_user)

	// updating movie status
	updatedMovieStatus, err := client.UpdateMovieStatus(ctx, &pb.Movie{Id: 3, Status: "Active"})
	utils.CheckError(err)
	fmt.Println(updatedMovieStatus)

	// updating show details
	updatedShowDetails, err := client.UpdateShowDetails(ctx, &pb.Show{Id: 1, MovieId: 2})
	utils.CheckError(err)
	fmt.Println(updatedShowDetails)

	//  deleting booking
	cancelledBooking, err := client.CancelBooking(ctx, &pb.Booking{Id: 3})
	utils.CheckError(err)
	fmt.Println(cancelledBooking)
}

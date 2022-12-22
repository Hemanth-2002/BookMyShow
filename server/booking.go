package server

import (
	pb "bms/bmsproto"
	model "bms/model"
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// function to add new booking on server
func (s *BmsServer) AddBooking(ctx context.Context, in *pb.NewBooking) (*pb.Booking, error) {
	log.Printf("creating new booking called")
	var wg sync.WaitGroup
	wg.Add(1)
	var newBooking = model.Booking{
		UserID: int(in.GetUserId()),
		ShowID: int(in.GetShowId()),
		Amount: int(in.GetAmount()),
	}
	go func(newBooking model.Booking) {
		defer wg.Done()
		time.Sleep(time.Duration(100) * time.Microsecond)
	}(newBooking)
	status, err := s.Db.AddBooking(newBooking)
	CheckCall(status, err)
	return &pb.Booking{UserId: in.GetUserId(), ShowId: in.GetShowId(), Amount: in.GetAmount(), Id: uint64(newBooking.ID)}, nil
}

// function to get list of bookings by user
func (s *BmsServer) GetListOfBookingsByUser(ctx context.Context, in *pb.User) (*pb.Bookings, error) {
	log.Printf("Getting list of bookings by user called")
	AllBookings := []*pb.Booking{}
	UserId := int(in.GetId())
	Bookings, status, err := s.Db.GetBookings(UserId)
	CheckCall(status, err)
	for _, booking := range Bookings {
		AllBookings = append(AllBookings, &pb.Booking{
			UserId: uint64(booking.UserID),
			ShowId: uint64(booking.ShowID),
			Amount: uint64(booking.Amount),
		})
	}
	return &pb.Bookings{Bookings: AllBookings}, nil
}

// function to cancel booking on server
func (s *BmsServer) CancelBooking(ctx context.Context, in *pb.Booking) (*pb.Booking, error) {
	log.Printf("cancel booking called")
	BookingId := in.Id
	status, err := s.Db.CancelBooking(int(BookingId))
	CheckCall(status, err)
	return &pb.Booking{Id: in.GetId()}, nil
}

func CheckCall(status bool, err error) {
	if status && err != nil {
		fmt.Println("call succesful")
	}
}

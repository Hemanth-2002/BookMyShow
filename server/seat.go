package server

import (
	pb "bms/bmsproto"
	model "bms/model"
	"bms/utils"
	"context"
	"log"
)

// function to add seats to booking on server
func (s *BmsServer) AddSeat(ctx context.Context, in *pb.NewSeat) (*pb.Seat, error) {
	log.Printf("adding seat called")
	newSeat := model.Seat{
		BookingID:  int(in.GetBookingId()),
		SeatNumber: int(in.GetSeatNumber()),
	}
	err := s.Db.AddSeat(newSeat)
	utils.CheckCall(err)
	if err != nil {
		return nil, err
	}
	return &pb.Seat{BookingId: uint64(newSeat.BookingID), SeatNumber: uint64(newSeat.SeatNumber), Id: uint64(newSeat.ID)}, nil
}

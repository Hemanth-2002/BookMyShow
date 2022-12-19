package server

import (
	pb "bms/bmsproto"
	model "bms/model"
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
	s.Db.Save(&newSeat)
	return &pb.Seat{BookingId: in.GetBookingId(), SeatNumber: in.GetSeatNumber(), Id: uint64(newSeat.ID)}, nil
}

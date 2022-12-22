package server

import (
	mail "bms/Mail"
	pb "bms/bmsproto"
	model "bms/model"
	"context"
	"fmt"
	"log"
)

// function to add payment for a booking
func (s *BmsServer) AddPayment(ctx context.Context, in *pb.NewPayment) (*pb.Payment, error) {
	log.Printf("creating new payment called")
	status := in.GetStatus()
	newPayment := model.Payment{
		BookingID:        int(in.GetBookingId()),
		Amount:           int(in.GetAmount()),
		DiscountCouponID: int(in.GetDiscountCouponId()),
		Mode:             in.GetMode(),
		Status:           status,
		UserID:           int(in.GetUserId()),
	}
	user, err := s.Db.AddPayment(newPayment)
	CheckCall(err)

	if status {
		sender := "hemanth.kakumanu@beautifulcode.in"
		mail.Mail(fmt.Sprint(sender), fmt.Sprint(in.GetAmount()), user.Email, fmt.Sprint(in.GetDiscountCouponId()), in.GetMode())
	}

	return &pb.Payment{Amount: in.GetAmount(), DiscountCouponId: in.GetDiscountCouponId(), Mode: in.GetMode(), Status: in.GetStatus(), Id: uint64(newPayment.ID)}, nil
}

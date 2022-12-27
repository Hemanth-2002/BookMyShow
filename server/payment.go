package server

import (
	mail "bms/Mail"
	pb "bms/bmsproto"
	model "bms/model"
	"bms/utils"
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
	utils.CheckCall(err)
	if err != nil {
		return nil, err
	}

	if status {
		sender := "hemanth.kakumanu@beautifulcode.in"
		jsonInfo := utils.MailInfo(fmt.Sprint(sender), fmt.Sprint(in.GetAmount()), user.Email, fmt.Sprint(in.GetDiscountCouponId()), in.GetMode())
		mail.Mail(jsonInfo)
	}

	return &pb.Payment{Amount: uint64(newPayment.Amount), DiscountCouponId: uint64(newPayment.DiscountCouponID), Mode: newPayment.Mode, Status: newPayment.Status, Id: uint64(newPayment.ID)}, nil
}

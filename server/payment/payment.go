package server

import (
	mail "bms/Mail"
	pb "bms/bmsproto"
	model "bms/model"
	"context"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type BmsServer struct {
	pb.UnimplementedBmsDatabaseCrudServer
	Db *gorm.DB
}

// function to add payment for a booking
func (s *BmsServer) AddPayment(ctx context.Context, in *pb.NewPayment) (*pb.Payment, error) {
	log.Printf("creating new payment called")
	status := in.GetStatus()
	NewPayment := model.Payment{
		BookingID:        int(in.GetBookingId()),
		Amount:           int(in.GetAmount()),
		DiscountCouponID: int(in.GetDiscountCouponId()),
		Mode:             in.GetMode(),
		Status:           status,
		UserID:           int(in.GetUserId()),
	}
	s.Db.Save(&NewPayment)
	user := &model.User{}
	s.Db.Where("id = ?", in.GetUserId()).Find(&user)

	if status {
		mail.Mail(fmt.Sprint(in.GetAmount()), user.Email, fmt.Sprint(in.GetDiscountCouponId()), in.GetMode())
	}

	return &pb.Payment{Amount: in.GetAmount(), DiscountCouponId: in.GetDiscountCouponId(), Mode: in.GetMode(), Status: in.GetStatus(), Id: uint64(NewPayment.ID)}, nil
}

package server

import (
	pb "bms/bmsproto"
	"bms/model"
	"context"
	"log"

	"github.com/jinzhu/gorm"
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

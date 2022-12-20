package server

import (
	pb "bms/bmsproto"
	"bms/model"
	"context"
	"log"
)

// function to add new user on server
func (s *BmsServer) CreateUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("creating new user called")
	newUser := model.User{
		UserName:    in.GetUserName(),
		Password:    in.GetPassword(),
		Email:       in.GetEmail(),
		PhoneNumber: in.GetPhoneNumber(),
	}
	s.Db.CreateUser(newUser)
	return &pb.User{UserName: in.GetUserName(), Password: in.GetPassword(), Email: in.GetEmail(), PhoneNumber: in.GetPhoneNumber(), Id: uint64(newUser.ID)}, nil
}

// function to update user info on server
func (s *BmsServer) UpdateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	log.Printf("update user info called")
	updatedUser := model.User{
		Password:    in.GetPassword(),
		Email:       in.GetEmail(),
		PhoneNumber: in.GetPhoneNumber(),
	}
	s.Db.UpdateUser(int(in.Id), updatedUser)
	return &pb.User{UserName: in.GetUserName(), Password: in.GetPassword(), Email: in.GetEmail(), PhoneNumber: in.GetPhoneNumber()}, nil
}

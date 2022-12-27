package server

import (
	pb "bms/bmsproto"
	"bms/model"
	"bms/utils"
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
	userId, err := s.Db.CreateUser(newUser)
	utils.CheckCall(err)
	if err != nil {
		return nil, err
	}
	return &pb.User{UserName: newUser.UserName, Password: newUser.Password, Email: newUser.Email, PhoneNumber: newUser.PhoneNumber, Id: uint64(userId)}, nil
}

// function to update user info on server
func (s *BmsServer) UpdateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	log.Printf("update user info called")
	updatedUser := model.User{
		Password:    in.GetPassword(),
		Email:       in.GetEmail(),
		PhoneNumber: in.GetPhoneNumber(),
	}
	updatedUser.ID = uint(in.Id)
	err := s.Db.UpdateUser(updatedUser)
	utils.CheckCall(err)
	if err != nil {
		return nil, err
	}
	return &pb.User{Password: updatedUser.Password, Email: updatedUser.Email, PhoneNumber: updatedUser.PhoneNumber}, nil
}

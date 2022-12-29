package server

import (
	pb "bms/bmsproto"
	"bms/utils"
	"context"
	"log"
	"time"
)

// function to create new token on server
func (s *BmsServer) CreateToken(ctx context.Context, in *pb.NewUser) (*pb.Token, error) {
	log.Printf("creating new token called")
	user, err := NewUserAuth(in.UserName, in.Email)
	utils.CheckError(err)
	jwtManager := JWTManager{
		SecretKey:     "secret",
		TokenDuration: 15 * time.Minute,
	}
	token, err := jwtManager.Generate(user)
	utils.CheckCall(err)
	if err != nil {
		return nil, err
	}
	return &pb.Token{Token: token}, nil
}

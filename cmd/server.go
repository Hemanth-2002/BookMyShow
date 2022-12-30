package main

import (
	auth "bms/authorization"
	pb "bms/bmsproto"
	"bms/database"
	"bms/model"
	"bms/server"
	"bms/utils"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	// port          = ":50051" // choosing port number
	secretKey     = "secret"
	port          = ":8080"
	tokenDuration = 15 * time.Minute
)

func runGRPCServer(
	jwtManager *auth.JWTManager,
	listener net.Listener,
	db *gorm.DB,
) error {
	interceptor := auth.NewAuthInterceptor(jwtManager)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptor.Unary()))
	pb.RegisterBmsDatabaseCrudServer(grpcServer, &server.BmsServer{
		Db: database.DBClient{Db: db},
	})

	reflection.Register(grpcServer)

	return grpcServer.Serve(listener)
}

func main() {

	// db connection
	fmt.Println("inside main")
	db, err := model.StartDB()
	utils.PanicError(err)

	listen, err := net.Listen("tcp", port)
	utils.PanicError(err)
	defer db.Close()

	log.Printf("Using port no %v", listen.Addr())

	jwtManager := auth.NewJWTManager(secretKey, tokenDuration)
	fmt.Println("jwt manager working")
	//create new server
	err = runGRPCServer(jwtManager, listen, db)
	log.Print(err)
	utils.CheckError(err)

}

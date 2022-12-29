package main

import (
	pb "bms/bmsproto"
	"bms/database"
	"bms/model"
	"bms/server"
	"bms/utils"
	"log"
	"net"
	"time"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port          = ":50051" // choosing port number
	secretKey     = "secret"
	tokenDuration = 15 * time.Minute
)

func runGRPCServer(
	jwtManager *server.JWTManager,
	listener net.Listener,
	db *gorm.DB,
) error {
	interceptor := server.NewAuthInterceptor(jwtManager)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptor.Unary()))
	pb.RegisterBmsDatabaseCrudServer(grpcServer, &server.BmsServer{
		Db: database.DBClient{Db: db},
	})

	reflection.Register(grpcServer)

	return grpcServer.Serve(listener)
}

func main() {

	// db connection
	db, err := model.StartDB()
	utils.PanicError(err)

	listen, err := net.Listen("tcp", port)
	utils.PanicError(err)
	defer db.Close()

	log.Printf("Using port no %v", listen.Addr())

	if err != nil {
		log.Fatal("cannot seed users: ", err)
	}

	jwtManager := server.NewJWTManager(secretKey, tokenDuration)

	//create new server
	err = runGRPCServer(jwtManager, listen, db)
	utils.CheckError(err)

}

package main

import (
	pb "bms/bmsproto"
	"bms/database"
	"bms/model"
	server "bms/server"
	"bms/utils"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051" // choosing port number
)

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("Unary Interceptor called: ", info.FullMethod)
	return handler(ctx, req)
}

func main() {

	// db connection
	db, err := model.StartDB()
	utils.PanicError(err)

	listen, err := net.Listen("tcp", port)
	utils.PanicError(err)
	defer db.Close()

	//create new server
	new_server := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
	)

	pb.RegisterBmsDatabaseCrudServer(new_server, &server.BmsServer{
		Db: database.DBClient{Db: db},
	})

	log.Printf("Using port no %v", listen.Addr())

	if err := new_server.Serve(listen); err != nil {
		log.Fatal(err.Error())
	}
}

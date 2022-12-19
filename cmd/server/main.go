package main

import (
	pb "bms/bmsproto"
	"bms/model"
	server "bms/server/payment"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

const (
	port = ":50051" // choosing port number
)

type BmsServer struct {
	pb.UnimplementedBmsDatabaseCrudServer
	Db *gorm.DB
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Println("Unary Interceptor called: ", info.FullMethod)
	return handler(ctx, req)
}

func main() {
	model.StartDB()
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err.Error())
	}

	//db connection
	db, err := gorm.Open("postgres", "user=postgres password=MohanNeelima@01 dbname=BookMyShow sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//create new server
	new_server := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
	)
	pb.RegisterBmsDatabaseCrudServer(new_server, &server.BmsServer{
		Db: db,
	})

	log.Printf("Using port no %v", listen.Addr())

	if err := new_server.Serve(listen); err != nil {
		log.Fatal(err.Error())
	}
}

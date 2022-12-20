package server

import (
	pb "bms/bmsproto"
	"bms/database"
)

type BmsServer struct {
	pb.UnimplementedBmsDatabaseCrudServer
	Db database.DataBase
}

package server

import (
	pb "bms/bmsproto"
	model "bms/model"
	"context"
	"log"
)

// function to get list of shows by theatre
func (s *BmsServer) GetListOfShowsByTheatre(ctx context.Context, in *pb.Theatre) (*pb.Shows, error) {
	log.Printf("Getting list of shows by theatre called")
	AllShows := []*pb.Show{}
	TheatreId := in.GetTheatreId()
	Shows := s.Db.GetShow(int(TheatreId))
	for _, show := range Shows {
		AllShows = append(AllShows, &pb.Show{
			Date:      show.Date,
			StartTime: show.StartTime,
			EndTime:   show.EndTime,
			MovieId:   uint64(show.MovieID),
			TheatreId: uint64(show.TheatreID),
		})
	}
	return &pb.Shows{Shows: AllShows}, nil
}

// function to update show details on server
func (s *BmsServer) UpdateShowDetails(ctx context.Context, in *pb.Show) (*pb.Show, error) {
	log.Printf("update show details called")
	Show := model.Show{
		MovieID: int(in.GetMovieId()),
	}
	s.Db.UpdateShow(int(in.Id), Show)
	return &pb.Show{Date: in.GetDate(), MovieId: in.GetMovieId()}, nil
}
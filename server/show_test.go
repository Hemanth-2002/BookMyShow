package server

import (
	pb "bms/bmsproto"
	"bms/mocks"
	model "bms/model"
	"bms/utils"
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

var MockShow = model.Show{
	Date:      "5-3-2022",
	StartTime: "18:00",
	EndTime:   "21:00",
}

var NewTheatre = pb.Theatre{TheatreId: 1}

var UpdateShow = model.Show{
	MovieID: 2,
}

var MockUpdateShow = pb.Show{
	MovieId: 2,
}

func TestUpdateShowDetails(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBase(controller)
	testShow := BmsServer{Db: mockDb}
	ctx := context.Background()
	UpdateShow.ID = 1
	mockDb.EXPECT().UpdateShow(UpdateShow).Return(nil)
	MockUpdateShow.Id = 1
	got, err := testShow.UpdateShowDetails(ctx, &MockUpdateShow)
	utils.CheckError(err)
	expected := &pb.Show{
		MovieId: 2,
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

func TestGetListOfShowsByTheatre(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBase(controller)
	testShow := BmsServer{Db: mockDb}
	ctx := context.Background()
	TheatreId := 1
	mockDb.EXPECT().GetShow(TheatreId).Return([]model.Show{MockShow}, nil)
	shows, err := testShow.GetListOfShowsByTheatre(ctx, &NewTheatre)
	got := shows.Shows
	utils.CheckError(err)
	expected := []*pb.Show{}
	expected = append(expected, &pb.Show{
		Date:      "5-3-2022",
		StartTime: "18:00",
		EndTime:   "21:00",
	})
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

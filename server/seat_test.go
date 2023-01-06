package server

import (
	pb "bms/bmsproto"
	"bms/mocks"
	"bms/model"
	"bms/utils"
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

var MockSeat = model.Seat{
	BookingID:  1,
	SeatNumber: 2,
}

var NewSeat = pb.NewSeat{
	BookingId:  1,
	SeatNumber: 2,
}

func TestAddSeat(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBase(controller)
	testSeat := BmsServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().AddSeat(MockSeat).Return(nil)
	got, err := testSeat.AddSeat(ctx, &NewSeat)
	utils.CheckError(err)
	expected := &pb.Seat{
		BookingId:  1,
		SeatNumber: 2,
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

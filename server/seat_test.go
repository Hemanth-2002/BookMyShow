package server

import (
	pb "bms/bmsproto"
	"bms/mocks"
	"bms/model"
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

var mockSeat = model.Seat{
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
	mockDb.EXPECT().AddSeat(mockSeat).Return(nil)
	got, err := testSeat.AddSeat(ctx, &NewSeat)
	model.CheckError(err)
	expected := &pb.Seat{
		BookingId:  1,
		SeatNumber: 2,
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

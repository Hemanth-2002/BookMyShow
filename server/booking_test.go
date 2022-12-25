package server

import (
	pb "bms/bmsproto"
	"bms/mocks"
	model "bms/model"
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

var MockBooking = model.Booking{
	UserID: 2,
	ShowID: 1,
	Amount: 150,
}

var NewBooking = pb.NewBooking{
	UserId: 2,
	ShowId: 1,
	Amount: 150,
}

var NewUser = pb.User{
	Id: 1,
}

var CancelledBooking = pb.Booking{}

func TestAddBooking(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBase(controller)
	testBooking := BmsServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().AddBooking(MockBooking).Return(nil)
	got, err := testBooking.AddBooking(ctx, &NewBooking)
	model.CheckError(err)
	expected := &pb.Booking{
		UserId: 2,
		ShowId: 1,
		Amount: 150,
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}

	// tests := []struct {
	// 	dbInput     model.Booking
	// 	serverInput pb.NewBooking
	// 	expected    *pb.Booking
	// }{
	// 	{dbInput: model.Booking{
	// 		UserID: 2, ShowID: 1, Amount: 150,
	// 	}, serverInput: pb.NewBooking{
	// 		UserId: 2, ShowId: 1, Amount: 150,
	// 	}, expected: &pb.Booking{
	// 		UserId: 2, ShowId: 1, Amount: 150,
	// 	}},
	// 	{dbInput: model.Booking{
	// 		UserID: 3, ShowID: 2, Amount: 1500,
	// 	}, serverInput: pb.NewBooking{
	// 		UserId: 3, ShowId: 2, Amount: 1500,
	// 	}, expected: &pb.Booking{
	// 		UserId: 3, ShowId: 2, Amount: 1500,
	// 	}},
	// }

	// for i, tc := range tests {
	// 	mockDb.EXPECT().AddBooking(tc.dbInput).Return(nil)
	// 	got, err := testBooking.AddBooking(ctx, &tc.serverInput)
	// 	model.CheckError(err)
	// 	if !reflect.DeepEqual(got, tc.expected) {
	// 		t.Errorf("The Function Retured is not expected one. got %v expected %v i = %v",
	// 			got, tc.expected, i+1)
	// 	}
	// }
}

func TestGetListOfBookingsByUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBase(controller)
	testBooking := BmsServer{Db: mockDb}
	ctx := context.Background()
	UserId := 1
	mockDb.EXPECT().GetBookings(UserId).Return([]model.Booking{MockBooking}, nil)
	bookings, err := testBooking.GetListOfBookingsByUser(ctx, &NewUser)
	got := bookings.Bookings
	model.CheckError(err)
	expected := []*pb.Booking{}
	expected = append(expected, &pb.Booking{
		UserId: 2,
		ShowId: 1,
		Amount: 150,
	})
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

func TestCancelBooking(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBase(controller)
	testBooking := BmsServer{Db: mockDb}
	ctx := context.Background()
	BookingId := 1
	mockDb.EXPECT().CancelBooking(BookingId).Return(nil)
	CancelledBooking.Id = 1
	got, err := testBooking.CancelBooking(ctx, &CancelledBooking)
	model.CheckError(err)
	expected := &pb.Booking{
		Id: 1,
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

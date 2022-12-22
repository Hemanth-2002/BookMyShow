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

var MockPayment = model.Payment{
	Amount:    150,
	Mode:      "credit-card",
	Status:    false,
	BookingID: 2,
	UserID:    1,
}

var NewPayment = pb.NewPayment{
	Amount:    150,
	Mode:      "credit-card",
	Status:    false,
	BookingId: 2,
	UserId:    1,
}

func TestAddPayment(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBase(controller)
	testPayment := BmsServer{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().AddPayment(MockPayment).Return(mockUser, nil)
	got, err := testPayment.AddPayment(ctx, &NewPayment)
	model.CheckError(err)
	expected := &pb.Payment{
		Amount: 150,
		Mode:   "credit-card",
		Status: false,
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

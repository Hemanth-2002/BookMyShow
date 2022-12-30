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

var mockUser = model.User{
	UserName:    "KHK",
	Password:    "123",
	Email:       "hemanth.kakumanu@beautifulcode.in",
	PhoneNumber: "9123",
}

var UserNew = pb.NewUser{
	UserName:    "KHK",
	Password:    "123",
	Email:       "hemanth.kakumanu@beautifulcode.in",
	PhoneNumber: "9123",
}

var UserUpdate = model.User{
	Password: "newPwd", PhoneNumber: "456", Email: "hemanthkakumanu1@gmail.com",
}

var MockUserUpdate = pb.User{
	Password: "newPwd", PhoneNumber: "456", Email: "hemanthkakumanu1@gmail.com",
}

func TestCreateUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBase(controller)
	testUser := BmsServer{Db: mockDb}
	ctx := context.Background()
	id := uint(1)
	mockDb.EXPECT().CreateUser(mockUser).Return(id, nil)
	got, err := testUser.CreateUser(ctx, &UserNew)
	utils.CheckError(err)
	expected := &pb.User{
		Id:          1,
		UserName:    "KHK",
		Password:    "123",
		Email:       "hemanth.kakumanu@beautifulcode.in",
		PhoneNumber: "9123",
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

func TestUpdateUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := mocks.NewMockDataBase(controller)
	testUser := BmsServer{Db: mockDb}
	ctx := context.Background()
	UserUpdate.ID = 1
	mockDb.EXPECT().UpdateUser(UserUpdate).Return(nil)
	MockUserUpdate.Id = 1
	got, err := testUser.UpdateUser(ctx, &MockUserUpdate)

	utils.CheckError(err)
	expected := &pb.User{
		Password: "newPwd", PhoneNumber: "456", Email: "hemanthkakumanu1@gmail.com",
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The Function Retured is not expected one. got %v expected %v",
			got, expected)
	}
}

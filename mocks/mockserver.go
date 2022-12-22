// Code generated by MockGen. DO NOT EDIT.
// Source: bms/database (interfaces: DataBase)

// Package mocks is a generated GoMock package.
package mocks

import (
	model "bms/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDataBase is a mock of DataBase interface.
type MockDataBase struct {
	ctrl     *gomock.Controller
	recorder *MockDataBaseMockRecorder
}

// MockDataBaseMockRecorder is the mock recorder for MockDataBase.
type MockDataBaseMockRecorder struct {
	mock *MockDataBase
}

// NewMockDataBase creates a new mock instance.
func NewMockDataBase(ctrl *gomock.Controller) *MockDataBase {
	mock := &MockDataBase{ctrl: ctrl}
	mock.recorder = &MockDataBaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataBase) EXPECT() *MockDataBaseMockRecorder {
	return m.recorder
}

// AddBooking mocks base method.
func (m *MockDataBase) AddBooking(arg0 model.Booking) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBooking", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBooking indicates an expected call of AddBooking.
func (mr *MockDataBaseMockRecorder) AddBooking(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBooking", reflect.TypeOf((*MockDataBase)(nil).AddBooking), arg0)
}

// AddMovie mocks base method.
func (m *MockDataBase) AddMovie(arg0 model.Movie) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMovie", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMovie indicates an expected call of AddMovie.
func (mr *MockDataBaseMockRecorder) AddMovie(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMovie", reflect.TypeOf((*MockDataBase)(nil).AddMovie), arg0)
}

// AddPayment mocks base method.
func (m *MockDataBase) AddPayment(arg0 model.Payment) (model.User, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPayment", arg0)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddPayment indicates an expected call of AddPayment.
func (mr *MockDataBaseMockRecorder) AddPayment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPayment", reflect.TypeOf((*MockDataBase)(nil).AddPayment), arg0)
}

// AddSeat mocks base method.
func (m *MockDataBase) AddSeat(arg0 model.Seat) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSeat", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSeat indicates an expected call of AddSeat.
func (mr *MockDataBaseMockRecorder) AddSeat(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSeat", reflect.TypeOf((*MockDataBase)(nil).AddSeat), arg0)
}

// CancelBooking mocks base method.
func (m *MockDataBase) CancelBooking(arg0 int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelBooking", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelBooking indicates an expected call of CancelBooking.
func (mr *MockDataBaseMockRecorder) CancelBooking(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelBooking", reflect.TypeOf((*MockDataBase)(nil).CancelBooking), arg0)
}

// CreateUser mocks base method.
func (m *MockDataBase) CreateUser(arg0 model.User) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockDataBaseMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockDataBase)(nil).CreateUser), arg0)
}

// GetBookings mocks base method.
func (m *MockDataBase) GetBookings(arg0 int) ([]model.Booking, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookings", arg0)
	ret0, _ := ret[0].([]model.Booking)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetBookings indicates an expected call of GetBookings.
func (mr *MockDataBaseMockRecorder) GetBookings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookings", reflect.TypeOf((*MockDataBase)(nil).GetBookings), arg0)
}

// GetMovie mocks base method.
func (m *MockDataBase) GetMovie(arg0 model.MoviePreference) ([]model.Movie, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovie", arg0)
	ret0, _ := ret[0].([]model.Movie)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMovie indicates an expected call of GetMovie.
func (mr *MockDataBaseMockRecorder) GetMovie(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovie", reflect.TypeOf((*MockDataBase)(nil).GetMovie), arg0)
}

// GetMovies mocks base method.
func (m *MockDataBase) GetMovies() ([]model.Movie, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovies")
	ret0, _ := ret[0].([]model.Movie)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMovies indicates an expected call of GetMovies.
func (mr *MockDataBaseMockRecorder) GetMovies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovies", reflect.TypeOf((*MockDataBase)(nil).GetMovies))
}

// GetShow mocks base method.
func (m *MockDataBase) GetShow(arg0 int) ([]model.Show, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShow", arg0)
	ret0, _ := ret[0].([]model.Show)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetShow indicates an expected call of GetShow.
func (mr *MockDataBaseMockRecorder) GetShow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShow", reflect.TypeOf((*MockDataBase)(nil).GetShow), arg0)
}

// UpdateMovie mocks base method.
func (m *MockDataBase) UpdateMovie(arg0 model.Movie) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMovie", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMovie indicates an expected call of UpdateMovie.
func (mr *MockDataBaseMockRecorder) UpdateMovie(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMovie", reflect.TypeOf((*MockDataBase)(nil).UpdateMovie), arg0)
}

// UpdateShow mocks base method.
func (m *MockDataBase) UpdateShow(arg0 model.Show) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateShow", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateShow indicates an expected call of UpdateShow.
func (mr *MockDataBaseMockRecorder) UpdateShow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateShow", reflect.TypeOf((*MockDataBase)(nil).UpdateShow), arg0)
}

// UpdateUser mocks base method.
func (m *MockDataBase) UpdateUser(arg0 model.User) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockDataBaseMockRecorder) UpdateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockDataBase)(nil).UpdateUser), arg0)
}

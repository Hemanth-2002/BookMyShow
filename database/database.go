package database

import (
	"bms/model"

	"github.com/jinzhu/gorm"
)

type DataBase interface {
	CreateUser(model.User) (bool, error)
	UpdateUser(model.User) (bool, error)
	GetShow(int) ([]model.Show, bool, error)
	UpdateShow(model.Show) (bool, error)
	AddSeat(model.Seat) (bool, error)
	AddPayment(model.Payment) (model.User, bool, error)
	AddMovie(model.Movie) (bool, error)
	GetMovies() ([]model.Movie, bool, error)
	GetMovie(model.MoviePreference) ([]model.Movie, bool, error)
	UpdateMovie(model.Movie) (bool, error)
	AddBooking(model.Booking) (bool, error)
	GetBookings(int) ([]model.Booking, bool, error)
	CancelBooking(int) (bool, error)
}

type DBClient struct {
	Db *gorm.DB
}

func (db DBClient) CreateUser(user model.User) (bool, error) {
	db.Db.Save(&user)
	return true, nil
}

func (db DBClient) UpdateUser(user model.User) (bool, error) {
	db.Db.Save(&user)
	return true, nil
}

func (db DBClient) GetShow(theatreId int) ([]model.Show, bool, error) {
	Shows := []model.Show{}
	db.Db.Where(&model.Show{TheatreID: int(theatreId)}).Find(&Shows)
	return Shows, true, nil
}

func (db DBClient) UpdateShow(show model.Show) (bool, error) {
	db.Db.Save(&show)
	return true, nil
}

func (db DBClient) AddSeat(seat model.Seat) (bool, error) {
	db.Db.Save(&seat)
	return true, nil
}

func (db DBClient) AddPayment(payment model.Payment) (model.User, bool, error) {
	user := model.User{}
	db.Db.Save(&payment)
	db.Db.Where("id = ?", payment.UserID).Find(&user)
	return user, true, nil
}

func (db DBClient) AddMovie(movie model.Movie) (bool, error) {
	db.Db.Save(&movie)
	return true, nil
}

func (db DBClient) GetMovies() ([]model.Movie, bool, error) {
	Movies := []model.Movie{}
	db.Db.Find(&Movies)
	return Movies, true, nil
}

func (db DBClient) GetMovie(movie model.MoviePreference) ([]model.Movie, bool, error) {
	Movies := []model.Movie{}
	db.Db.Where(movie).Find(&Movies)
	return Movies, true, nil
}

func (db DBClient) UpdateMovie(movie model.Movie) (bool, error) {
	db.Db.Save(&movie)
	return true, nil
}

func (db DBClient) AddBooking(booking model.Booking) (bool, error) {
	db.Db.Save(&booking)
	return true, nil
}

func (db DBClient) GetBookings(userId int) ([]model.Booking, bool, error) {
	Bookings := []model.Booking{}
	db.Db.Where("user_id = ?", userId).Find(&Bookings)
	return Bookings, true, nil
}

func (db DBClient) CancelBooking(bookingId int) (bool, error) {
	db.Db.Model(&model.Booking{}).Where("id=?", bookingId).Delete(&model.Booking{})
	return true, nil
}

package database

import (
	"bms/model"

	"github.com/jinzhu/gorm"
)

type DataBase interface {
	CreateUser(model.User)
	UpdateUser(model.User)
	GetShow(int) []model.Show
	UpdateShow(model.Show)
	AddSeat(model.Seat)
	AddPayment(model.Payment) model.User
	AddMovie(model.Movie)
	GetMovies() []model.Movie
	GetMovie(model.MoviePreference) []model.Movie
	UpdateMovie(model.Movie)
	AddBooking(model.Booking)
	GetBookings(int) []model.Booking
	CancelBooking(int)
}

type DBClient struct {
	Db *gorm.DB
}

func (db DBClient) CreateUser(user model.User) {
	db.Db.Save(&user)
}

func (db DBClient) UpdateUser(user model.User) {
	db.Db.Save(&user)
}

func (db DBClient) GetShow(theatreId int) []model.Show {
	Shows := []model.Show{}
	db.Db.Where(&model.Show{TheatreID: int(theatreId)}).Find(&Shows)
	return Shows
}

func (db DBClient) UpdateShow(show model.Show) {
	db.Db.Save(&show)
}

func (db DBClient) AddSeat(seat model.Seat) {
	db.Db.Save(&seat)
}

func (db DBClient) AddPayment(payment model.Payment) model.User {
	user := model.User{}
	db.Db.Save(&payment)
	db.Db.Where("id = ?", payment.UserID).Find(&user)
	return user
}

func (db DBClient) AddMovie(movie model.Movie) {
	db.Db.Save(&movie)
}

func (db DBClient) GetMovies() []model.Movie {
	Movies := []model.Movie{}
	db.Db.Find(&Movies)
	return Movies
}

func (db DBClient) GetMovie(movie model.MoviePreference) []model.Movie {
	Movies := []model.Movie{}
	db.Db.Where(movie).Find(&Movies)
	return Movies
}

func (db DBClient) UpdateMovie(movie model.Movie) {
	db.Db.Save(&movie)
}

func (db DBClient) AddBooking(booking model.Booking) {
	db.Db.Save(&booking)
}

func (db DBClient) GetBookings(userId int) []model.Booking {
	Bookings := []model.Booking{}
	db.Db.Where("id = ?", userId).Find(&Bookings)
	return Bookings
}

func (db DBClient) CancelBooking(bookingId int) {
	db.Db.Model(&model.Booking{}).Where("id=?", bookingId).Delete(&model.Booking{})
}

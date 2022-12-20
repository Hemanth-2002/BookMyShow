package database

import (
	"bms/model"

	"github.com/jinzhu/gorm"
)

type DataBase interface {
	CreateUser(model.User)
	UpdateUser(int, model.User)
	GetShow(int) []model.Show
	UpdateShow(int, model.Show)
	AddSeat(model.Seat)
	AddPayment(model.Payment, int) model.User
	AddMovie(model.Movie)
	GetMovies() []model.Movie
	GetMovie(model.Movie) []model.Movie
	UpdateMovie(int, model.Movie)
	AddBooking(model.Booking)
	GetBookings(model.Booking) []model.Booking
	CancelBooking(int)
}

type DBClient struct {
	Db *gorm.DB
}

func (db DBClient) CreateUser(user model.User) {
	db.Db.Save(&user)
}

func (db DBClient) UpdateUser(userID int, user model.User) {
	db.Db.Model(&model.User{}).Where("id=?", userID).Updates(&user)
}

func (db DBClient) GetShow(theatreId int) []model.Show {
	Shows := []model.Show{}
	db.Db.Where(&model.Show{TheatreID: int(theatreId)}).Find(&Shows)
	return Shows
}

func (db DBClient) UpdateShow(showId int, show model.Show) {
	db.Db.Model(&model.User{}).Where("id=?", showId).Updates(&show)
}

func (db DBClient) AddSeat(seat model.Seat) {
	db.Db.Save(&seat)
}

func (db DBClient) AddPayment(payment model.Payment, userId int) model.User {
	user := &model.User{}
	db.Db.Save(&payment)
	db.Db.Where("id = ?", userId).Find(&user)
	return *user
}

func (db DBClient) AddMovie(movie model.Movie) {
	db.Db.Save(&movie)
}

func (db DBClient) GetMovies() []model.Movie {
	Movies := []model.Movie{}
	db.Db.Find(&Movies)
	return Movies
}

func (db DBClient) GetMovie(movie model.Movie) []model.Movie {
	Movies := []model.Movie{}
	db.Db.Where(movie).Find(&Movies)
	return Movies
}

func (db DBClient) UpdateMovie(movieId int, movie model.Movie) {
	db.Db.Model(&model.User{}).Where("id=?", movieId).Updates(&movie)
}

func (db DBClient) AddBooking(booking model.Booking) {
	db.Db.Save(&booking)
}

func (db DBClient) GetBookings(booking model.Booking) []model.Booking {
	Bookings := []model.Booking{}
	db.Db.Where(booking).Find(&Bookings)
	return Bookings
}

func (db DBClient) CancelBooking(userId int) {
	db.Db.Model(&model.Booking{}).Where("id=?", userId).Delete(&model.Booking{})
}

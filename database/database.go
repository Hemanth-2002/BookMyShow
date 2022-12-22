package database

import (
	"bms/model"

	"github.com/jinzhu/gorm"
)

type DataBase interface {
	CreateUser(model.User) error
	UpdateUser(model.User) error
	GetShow(int) ([]model.Show, error)
	UpdateShow(model.Show) error
	AddSeat(model.Seat) error
	AddPayment(model.Payment) (model.User, error)
	AddMovie(model.Movie) error
	GetMovies() ([]model.Movie, error)
	GetMovie(model.MoviePreference) ([]model.Movie, error)
	UpdateMovie(model.Movie) error
	AddBooking(model.Booking) error
	GetBookings(int) ([]model.Booking, error)
	CancelBooking(int) error
}

type DBClient struct {
	Db *gorm.DB
}

func (db DBClient) CreateUser(user model.User) error {
	DB := db.Db.Save(&user)
	return DB.Error
}

func (db DBClient) UpdateUser(user model.User) error {
	DB := db.Db.Save(&user)
	return DB.Error
}

func (db DBClient) GetShow(theatreId int) ([]model.Show, error) {
	Shows := []model.Show{}
	DB := db.Db.Where(&model.Show{TheatreID: int(theatreId)}).Find(&Shows)
	return Shows, DB.Error
}

func (db DBClient) UpdateShow(show model.Show) error {
	DB := db.Db.Save(&show)
	return DB.Error
}

func (db DBClient) AddSeat(seat model.Seat) error {
	DB := db.Db.Save(&seat)
	return DB.Error
}

func (db DBClient) AddPayment(payment model.Payment) (model.User, error) {
	user := model.User{}
	db.Db.Save(&payment)
	DB := db.Db.Where("id = ?", payment.UserID).Find(&user)
	return user, DB.Error
}

func (db DBClient) AddMovie(movie model.Movie) error {
	DB := db.Db.Save(&movie)
	return DB.Error
}

func (db DBClient) GetMovies() ([]model.Movie, error) {
	Movies := []model.Movie{}
	DB := db.Db.Find(&Movies)
	return Movies, DB.Error
}

func (db DBClient) GetMovie(movie model.MoviePreference) ([]model.Movie, error) {
	Movies := []model.Movie{}
	DB := db.Db.Where(movie).Find(&Movies)
	return Movies, DB.Error
}

func (db DBClient) UpdateMovie(movie model.Movie) error {
	DB := db.Db.Save(&movie)
	return DB.Error
}

func (db DBClient) AddBooking(booking model.Booking) error {
	DB := db.Db.Save(&booking)
	return DB.Error
}

func (db DBClient) GetBookings(userId int) ([]model.Booking, error) {
	Bookings := []model.Booking{}
	DB := db.Db.Where("user_id = ?", userId).Find(&Bookings)
	return Bookings, DB.Error
}

func (db DBClient) CancelBooking(bookingId int) error {
	DB := db.Db.Model(&model.Booking{}).Where("id=?", bookingId).Delete(&model.Booking{})
	return DB.Error
}

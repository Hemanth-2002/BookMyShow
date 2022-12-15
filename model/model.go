// package main
package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func StartDB() {
	// func main() {

	db, err := gorm.Open("postgres", "user=postgres password=MohanNeelima@01 dbname=BookMyShow sslmode=disable")

	CheckError(err)
	defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Booking{})
	db.AutoMigrate(&Movie{})
	db.AutoMigrate(&Show{})
	db.AutoMigrate(&Theatre{})
	db.AutoMigrate(&Seat{})
	db.AutoMigrate(&Payment{})

	db.Model(&Booking{}).AddForeignKey("user_id", "users(id)", "CASCADE", "RESTRICT")
	db.Model(&Booking{}).AddForeignKey("show_id", "shows(id)", "CASCADE", "RESTRICT")
	db.Model(&Payment{}).AddForeignKey("booking_id", "bookings(id)", "CASCADE", "RESTRICT")
	db.Model(&Seat{}).AddForeignKey("booking_id", "bookings(id)", "CASCADE", "RESTRICT")
	db.Model(&Show{}).AddForeignKey("movie_id", "movies(id)", "CASCADE", "RESTRICT")
	db.Model(&Show{}).AddForeignKey("theatre_id", "theatres(id)", "CASCADE", "RESTRICT")

	// db.Save(&User{
	// 	UserName:    "KHK",
	// 	Password:    "123",
	// 	Email:       "khk@bc.in",
	// 	PhoneNumber: "100",
	// })

	// db.Save(&Movie{
	// 	Model:       gorm.Model{},
	// 	MovieName:   "RRR",
	// 	Director:    "Rajamouli",
	// 	Description: "...plot of the movie...",
	// 	Rating:      8,
	// 	Language:    "Telugu",
	// 	Genre:       "Action",
	// 	ReleaseDate: "1-3-2022",
	// 	Status:      "Active",
	// })

	// db.Save(&Theatre{
	// 	Model:       gorm.Model{},
	// 	Location:    "Nagole",
	// 	TheatreName: "PVR",
	// 	TotalSeats:  200,
	// })

	// db.Save(&Show{
	// 	Date: "5-3-2022", StartTime: "11:00", EndTime: "13:00", MovieID: 2, TheatreID: 1,
	// })

	// db.Save(&Booking{
	// 	UserID: 1,
	// 	Amount: 100,
	// 	ShowID: 1,
	// })

	// db.Save(&Payment{
	// 	Amount:           100,
	// 	DiscountCouponID: 123,
	// 	Mode:             "Upi",
	// 	Status:           true,
	// 	BookingID:        1,
	// })

	// db.Save(&Seat{
	// 	SeatNumber: 11,
	// 	BookingID:  1,
	// })

	// db.Save(&Seat{
	// 	SeatNumber: 12,
	// 	BookingID:  1,
	// })

	// db.Save(&Seat{
	// 	SeatNumber: 13,
	// 	BookingID:  1,
	// })

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

type User struct {
	gorm.Model
	UserName    string
	Password    string
	Email       string
	PhoneNumber string
	Booking     []Booking
}

type Booking struct {
	gorm.Model
	UserID  int
	ShowID  int
	Amount  int
	Payment []Payment
	Seat    []Seat
}

type Movie struct {
	gorm.Model
	MovieName   string
	Director    string
	Description string
	Rating      int
	Language    string
	Genre       string
	ReleaseDate string
	Status      string
	Show        []Show
}

type Show struct {
	gorm.Model
	Date      string
	StartTime string
	EndTime   string
	MovieID   int
	TheatreID int
	Booking   []Booking
}

type Theatre struct {
	gorm.Model
	Location    string
	TheatreName string
	TotalSeats  int
	Show        []Show
}

type Seat struct {
	gorm.Model
	SeatNumber int
	BookingID  int
}

type Payment struct {
	gorm.Model
	Amount           int
	DiscountCouponID int
	Mode             string
	Status           bool
	BookingID        int
}

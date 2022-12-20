// package main
package model

import "github.com/jinzhu/gorm"

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

type MoviePreference struct {
	Language string
	Rating   int
	Genre    string
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
	UserID           int
}

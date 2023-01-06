package model

import "github.com/jinzhu/gorm"

func initializeDb(db *gorm.DB) {

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
	// 	UserID:           1,
	// })

	// db.Save(&Payment{
	// 	Amount:    150,
	// 	Mode:      "Credit-card",
	// 	Status:    true,
	// 	BookingID: 2,
	// 	UserID:    2,
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

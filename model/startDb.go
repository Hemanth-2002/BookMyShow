package model

import (
	"flag"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func StartDB() (*gorm.DB, error) {

	db_user := flag.String("user", "postgres", "database user")
	db_password := flag.String("password", "MohanNeelima@01", "database password")
	db_name := flag.String("name", "BookMyShow", "database name")

	conn := "user=" + *db_user + " password=" + *db_password + " dbname=" + *db_name + " sslmode=disable"
	db, err := gorm.Open("postgres", conn)
	CheckError(err)

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

	return db, nil
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

package model

import (
	"bms/utils"
	"flag"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func StartDB() (*gorm.DB, error) {

	utils.LoadEnv() // To load .env file

	db_user := flag.String("user", "postgres", "database user")
	db_password := os.Getenv("db_password")
	db_name := os.Getenv("db_name")
	conn := "user=" + *db_user + " password=" + db_password + " dbname=" + db_name + " sslmode=disable"
	db, err := gorm.Open("postgres", conn)
	utils.PanicError(err)

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

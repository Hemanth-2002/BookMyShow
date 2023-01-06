package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckError(err error) {
	if err != nil {
		log.Print(err)
	}
}

func CheckCall(err error) {
	if err != nil {
		fmt.Println("call unsuccesfull")
	}
}

func LoadEnv() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("Could not load .env file")
		os.Exit(1)
	}
}

func MailInfo(sender string, amount string, receiver string, coupon string, mode string) string {
	jsonInfo := fmt.Sprintf(`{
	    "personalizations": [
	        {
	            "to": [
	                {
	                    "email": `+`"%s"`+`
	                }
	            ],
	            "subject": "Payment Report"
	        }
	    ],
	    "from": {
	        "email": `+`"%s"`+`
	    },
	    "content": [
	        {
	            "type": "text/plain",
	            "value":`+` "Your payment is successful !!!\n\nThe following are the payment details:\n\nAmount: `+fmt.Sprint(amount)+`\n\nDiscount Coupon Number: `+fmt.Sprint(coupon)+`\n\nPayment Mode: `+fmt.Sprint(mode)+`\n\nThank You For Booking With Us"`+`
			}
	    ]
	}`, receiver, sender)
	return jsonInfo
}

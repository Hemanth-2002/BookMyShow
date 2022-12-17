package Mail

import (
	model "bms/model"
	"fmt"
	"net/http"
	"strings"
)

// func main() {
func Mail(amount string, email string, coupon string, mode string) {
	url := "https://rapidprod-sendgrid-v1.p.rapidapi.com/mail/send"
	code := fmt.Sprintf(`{
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
	        "email": "hemanth.kakumanu@beautifulcode.in"
	    },
	    "content": [
	        {
	            "type": "text/plain",
	            "value":`+` "Your payment is successful !!!\n\nThe following are the payment details:\n\nAmount: `+fmt.Sprint(amount)+`\n\nDiscount Coupon Number: `+fmt.Sprint(coupon)+`\n\nPayment Mode: `+fmt.Sprint(mode)+`\n\nThank You For Booking With Us"`+`
			}
	    ]
	}`, email)
	fmt.Println(code)

	payload := strings.NewReader(code)
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", "81087b3865mshb8974e3c74f5b9dp1999ffjsn823050f51b45")
	req.Header.Add("X-RapidAPI-Host", "rapidprod-sendgrid-v1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	model.CheckError(err)
	defer res.Body.Close()
	// body, _ := io.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))

}

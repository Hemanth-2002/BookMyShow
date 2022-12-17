package Mail

import (
	"bms/model"
	"fmt"
	"net/http"
	"strings"
)

// func main() {
func Mail(amount string, email string, coupon string, mode string) {
	url := "https://rapidprod-sendgrid-v1.p.rapidapi.com/mail/send"
	fmt.Println(email)
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
	            "value":`+` "Your payment is successful !!!\nThe following are the payment details\nAmount: `+fmt.Sprint(amount)+`\nDiscountCouponId: `+fmt.Sprint(coupon)+`\nMode: `+`%s"`+`
			}
	    ]
	}`, email, mode)

	payload := strings.NewReader(code)
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", "285bd11398msh3a5f0f9c659453fp128fd7jsn73c2bb09d864")
	req.Header.Add("X-RapidAPI-Host", "rapidprod-sendgrid-v1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	model.CheckError(err)
	defer res.Body.Close()

}

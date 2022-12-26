package Mail

import (
	"bms/utils"
	"net/http"
	"os"
	"strings"
)

// func main() {
func Mail(sender string, amount string, receiver string, coupon string, mode string) {
	url := "https://rapidprod-sendgrid-v1.p.rapidapi.com/mail/send"

	jsonInfo := utils.MailInfo(sender, amount, receiver, coupon, mode)

	payload := strings.NewReader(jsonInfo)
	req, _ := http.NewRequest("POST", url, payload)

	key := os.Getenv("X-RapidAPI-Key")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", key)
	req.Header.Add("X-RapidAPI-Host", "rapidprod-sendgrid-v1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	utils.CheckError(err)
	defer res.Body.Close()

}

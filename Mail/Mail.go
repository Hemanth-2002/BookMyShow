package Mail

import (
	"bms/utils"
	"net/http"
	"strings"
)

// func main() {
func Mail(jsonInfo string) {
	url := "https://rapidprod-sendgrid-v1.p.rapidapi.com/mail/send"

	payload := strings.NewReader(jsonInfo)
	req, _ := http.NewRequest("POST", url, payload)

	// key := os.Getenv("X-RapidAPI-Key")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", "81087b3865mshb8974e3c74f5b9dp1999ffjsn823050f51b45")
	req.Header.Add("X-RapidAPI-Host", "rapidprod-sendgrid-v1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	utils.CheckError(err)
	defer res.Body.Close()

}

package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig *oauth2.Config

func init() {
	LoadEnv()
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  os.Getenv("GOOGLE_OAUTH_REDIRECT_URL"),
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

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
		fmt.Println("call unsuccessful")
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
	                    "email": "%s"
	                }
	            ],
	            "subject": "Payment Report"
	        }
	    ],
	    "from": {
	        "email": "%s"
	    },
	    "content": [
	        {
	            "type": "text/plain",
	            "value": "Your payment is successful !!!\n\nThe following are the payment details:\n\nAmount: %s\n\nDiscount Coupon Number: %s\n\nPayment Mode: %s\n\nThank You For Booking With Us"
			}
	    ]
	}`, receiver, sender, amount, coupon, mode)
	return jsonInfo
}

func GenerateGoogleOAuthURL(state string) string {
	return googleOauthConfig.AuthCodeURL(state)
}

func ExchangeCodeForTokens(code string) (*oauth2.Token, error) {
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code for token: %v", err)
	}
	return token, nil
}

func GetGoogleUserInfo(token *oauth2.Token) (map[string]interface{}, error) {
	client := googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %v", err)
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %v", err)
	}
	return userInfo, nil
}

func HandleOAuth2Error(w http.ResponseWriter, err error) {
	http.Error(w, fmt.Sprintf("OAuth2 error: %v", err), http.StatusInternalServerError)
}
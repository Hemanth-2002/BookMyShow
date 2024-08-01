package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	pb "bms/bmsproto"
	"bms/model"
	"bms/utils"
)

// Google OAuth2 configuration
var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/auth/google/callback",
	ClientID:     "your-client-id",
	ClientSecret: "your-client-secret",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}

// function to add new user on server
func (s *BmsServer) CreateUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("creating new user called")
	newUser := model.User{
		UserName:    in.GetUserName(),
		Password:    in.GetPassword(),
		Email:       in.GetEmail(),
		PhoneNumber: in.GetPhoneNumber(),
	}
	userId, err := s.Db.CreateUser(newUser)
	utils.CheckCall(err)
	if err != nil {
		return nil, err
	}
	return &pb.User{UserName: newUser.UserName, Password: newUser.Password, Email: newUser.Email, PhoneNumber: newUser.PhoneNumber, Id: uint64(userId)}, nil
}

// function to update user info on server
func (s *BmsServer) UpdateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	log.Printf("update user info called")
	updatedUser := model.User{
		Password:    in.GetPassword(),
		Email:       in.GetEmail(),
		PhoneNumber: in.GetPhoneNumber(),
	}
	updatedUser.ID = uint(in.Id)
	err := s.Db.UpdateUser(updatedUser)
	utils.CheckCall(err)
	if err != nil {
		return nil, err
	}
	return &pb.User{Password: updatedUser.Password, Email: updatedUser.Email, PhoneNumber: updatedUser.PhoneNumber}, nil
}

// Handler to initiate Google OAuth2 sign-in process
func (s *BmsServer) HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Handler to handle Google OAuth2 callback
func (s *BmsServer) HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != "state-token" {
		http.Error(w, "State token mismatch", http.StatusBadRequest)
		return
	}

	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	client := googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		http.Error(w, "Failed to decode user info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	userInfoJSON, err := json.Marshal(userInfo)
	if err != nil {
		http.Error(w, "Failed to marshal user info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(userInfoJSON)
}
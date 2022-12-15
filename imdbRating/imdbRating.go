package imdbRating

// package main

import (
	"bms/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ImdbRating(movieName string) int {

	url := fmt.Sprintf("https://mdblist.p.rapidapi.com/?s=%s&m=movie&l=1", movieName)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "285bd11398msh3a5f0f9c659453fp128fd7jsn73c2bb09d864")
	req.Header.Add("X-RapidAPI-Host", "mdblist.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	model.CheckError(err)
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var responseObj Response
	json.Unmarshal(body, &responseObj)

	rating := responseObj.MovieRating[0].AvgScore / 10
	return rating
}

type Response struct {
	MovieRating []MovieRating `json:"search"`
}

type MovieRating struct {
	AvgScore int `json:"score_average"`
}

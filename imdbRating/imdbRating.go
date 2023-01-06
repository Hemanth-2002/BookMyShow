package imdbRating

import (
	"bms/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetImdbRating(movieName string) int {

	url := fmt.Sprintf("https://mdblist.p.rapidapi.com/?s=%s&m=movie&l=1", movieName)

	req, _ := http.NewRequest("GET", url, nil)
	key := os.Getenv("RapidAPI_Key_Rating")
	req.Header.Add("X-RapidAPI-Key", key)
	req.Header.Add("X-RapidAPI-Host", "mdblist.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	utils.CheckError(err)
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

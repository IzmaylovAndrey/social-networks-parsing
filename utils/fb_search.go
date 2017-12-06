package utils

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"net/url"
	"encoding/json"
)

type FBData struct {
	FBData []Person `json:"data"`
	Paging Link `json:"paging"`
}

type Person struct {
	Name string `json:"name"`
	ID string `json:"id"`
}

type Link struct {
	Next string `json:"next"`
}

func FBSearch(name string) (result []string, err error) {
	token := "EAACEdEose0cBALju0qCsGXkVUyaHtcrnfzoFbBWi1D60GFwKMYogxIPbUhBCJM2BIJSECk1ZBQOTkt7fLH1tqlqZCICAGAJVFVZC8AqLlzqRtsgEKQ1Q22IWTZCJd4ZCvYEkBFQpOLTPp6y8WN0fWxdgkGEnvI7AKO5I9EW6IZBbGZCuiF9zJhlMxULmuCZBcRZACnztG7QYPOgZDZD"

	response, err := http.Get("https://graph.facebook.com/search?q=" + url.QueryEscape(name) + "&type=user&limit=3&access_token=" + token)
	if err != nil{
		fmt.Printf("Error sending request to FB: %s", err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		var data FBData

		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil{
			fmt.Printf("Error getting answer from FB: %s", err)
			return nil, err
		}

		err = json.Unmarshal(bodyBytes, &data)
		if err != nil{
			fmt.Printf("Error unmarshaling FB answer: %s", err)
			return nil, err
		}

		for i := 0; i < len(data.FBData); i++ {
			result = append(result, "https://facebook.com/" + data.FBData[i].ID)
		}
	}
	return
}

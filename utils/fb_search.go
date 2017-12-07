package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

type FBData struct {
	FBData []Person `json:"data"`
	Paging Link     `json:"paging"`
}

type Person struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type Link struct {
	Next string `json:"next"`
}

func FBSearch(name string, commonResult *APIHandlersResult, wg *sync.WaitGroup) {
	defer wg.Done()
	token := "EAACEdEose0cBAK7JMZA1rGkiZCUMeDWflA0A6WZAEUX72ZC7acnTA4JDb3ZCyOChvuGVrCXexn83drUsZBZAZBmrzNvjBzcUqQwbWX6bbbcbh6rgCsbXhggyDHFR24XaxnuDH3BknCgZAPrOPkR3Yme0lrluHZBe3eFRf4tAQG4FJ7ys13hp1ASy02VeRtmmMghEZCphfkKZAZAOvFQZDZD"

	response, err := http.Get("https://graph.facebook.com/search?q=" + url.QueryEscape(name) + "&type=user&limit=3&access_token=" + token)
	if err != nil {
		fmt.Printf("Error sending request to FB: %s", err)
		commonResult.Lock()
		commonResult.Errors = append(commonResult.Errors, err)
		commonResult.Unlock()
		return
	}
	defer response.Body.Close()

	var result []string
	if response.StatusCode == http.StatusOK {
		var data FBData

		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Error getting answer from FB: %s", err)
			commonResult.Lock()
			commonResult.Errors = append(commonResult.Errors, err)
			commonResult.Unlock()
		}

		err = json.Unmarshal(bodyBytes, &data)
		if err != nil {
			fmt.Printf("Error unmarshaling FB answer: %s", err)
			commonResult.Lock()
			commonResult.Errors = append(commonResult.Errors, err)
			commonResult.Unlock()
		}

		for _, v := range data.FBData {
			result = append(result, "https://facebook.com/"+v.ID)
		}
	}
	commonResult.Lock()
	commonResult.Facebook = result
	commonResult.Unlock()

}

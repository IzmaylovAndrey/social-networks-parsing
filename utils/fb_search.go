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
	token := "EAACEdEose0cBADlVArrgPz86FYumZB44zgBFTBpsFVIxZAMahjHtSt2aVmpbUxzbCkzddLTYKcXkdbkvmH5wIrtgFKs4VfS7K1KsE3hCMA5n0C48STz2x5mKV3rfZB8vtcKsRN6x3wahazQRWF9HZB0HjlqcCaBSNzZAnuACStHJgXlgcwAmYZAMMN6TOV470aPVGjANicogZDZD"

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

		for _, v := range data.FBData {
			result = append(result, "https://facebook.com/" + v.ID)
		}
	}
	return
}

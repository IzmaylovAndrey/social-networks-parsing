package utils

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"net/url"
	"encoding/json"
)

type Data struct {
	Data []Person `json:"data"`
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
	token := "EAACEdEose0cBAJupgdWdBrHtFPkBdI1RQDaJ7MdPeZCi7ZCNVZAGgxGGmiIOdrBkSqZCb58VZAvuuBnLt5s6QAzzVqaPzSOCbYKRAXCjnHAvvsT3iKshKz4vStORGBPWFNZC9a6Dt6xEw6hnRxpuM3mzLHVaZAZAZBouYjGyZBOhdLr9yb8wUga03YHVwlC9ghRfkcFs3pxKSA1QZDZD"

	response, err := http.Get("https://graph.facebook.com/search?q=" + url.QueryEscape(name) + "&type=user&limit=3&access_token=" + token)
	if err != nil{
		fmt.Printf("Error sending request to FB: %s", err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		var data Data

		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil{
			fmt.Printf("Error getting answer from FB: %s", err)
			return nil, err
		}

		err = json.Unmarshal(bodyBytes, &data)
		if err != nil{
			fmt.Printf("Error unmarshaling: %s", err)
			return nil, err
		}

		for i := 0; i < len(data.Data); i++ {
			result = append(result, "https://facebook.com/" + data.Data[i].ID)
		}
	}
	return
}

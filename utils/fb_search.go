package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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

func FBSearch(name string, ch chan ChanelResult) {
	token := "EAACEdEose0cBAJ7WWCgLqWBui98vZAeuqQ3qfbZB1JfZAEng4jv3gvuQ0pI5lygMKTIwfErdDIhpCYBz6CGGJsP8MMi0WwKUboya8yH26ZCZAXBPfMjUTLcy3wzcZCkc9VD56HpC7vBCKCeXHGOvwc4dJD6PKWbzcKjdZB81fC9ZBZBAKeYGmJqTpjfYDIUdED3ugUZAocatHYLgZDZD"

	response, err := http.Get("https://graph.facebook.com/search?q=" + url.QueryEscape(name) + "&type=user&limit=3&access_token=" + token)
	if err != nil {
		fmt.Printf("Error sending request to FB: %s", err)
		ch <- ChanelResult{nil, err}
	}
	defer response.Body.Close()

	result := make([]string, 3)
	if response.StatusCode == http.StatusOK {
		var data FBData

		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Error getting answer from FB: %s", err)
			ch <- ChanelResult{nil, err}
		}

		err = json.Unmarshal(bodyBytes, &data)
		if err != nil {
			fmt.Printf("Error unmarshaling FB answer: %s", err)
			ch <- ChanelResult{nil, err}
		}

		for _, v := range data.FBData {
			result = append(result, "https://facebook.com/"+v.ID)
		}
	}
	ch <- ChanelResult{result, nil}
}

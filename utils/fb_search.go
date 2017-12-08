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
	token := "EAACEdEose0cBABe9ScW8SSDtX9vKJJ1EH2dGc93OYsTkinGnUE3OyxlfVn9lrZC3y8zgCb7ZC0slZAjgJpVSrRKlG76ZBteTpUfZBkmYdKAVSysRBZBt7gnQjibxmrdPMvnM1yMOMwug1hWP3abfKan7dsECRk98kMrkfOSKKIKowfIzz4KsMHEtyo0N6qQr1u3DJSEHZBXhAZDZD"

	response, err := http.Get("https://graph.facebook.com/search?q=" + url.QueryEscape(name) + "&type=user&limit=5&access_token=" + token)
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
			return
		}

		err = json.Unmarshal(bodyBytes, &data)
		if err != nil {
			fmt.Printf("Error unmarshaling FB answer: %s", err)
			commonResult.Lock()
			commonResult.Errors = append(commonResult.Errors, err)
			commonResult.Unlock()
			return
		}

		for _, v := range data.FBData {
			result = append(result, "https://facebook.com/"+v.ID)
		}
	}
	commonResult.Lock()
	commonResult.Facebook = result
	commonResult.Unlock()
}

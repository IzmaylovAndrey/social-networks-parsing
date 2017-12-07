package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type VKData struct {
	Response []json.RawMessage `json:"response"`
}

type VKPerson struct {
	UID       int    `json:"uid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func VKSearch(name string, ch chan ChanelResult) {
	token := "2e11eacee8ccc36fd2e2d982672e95900e08ed6fff29037ec475b8847ec8ce138305b94df5b95100399b2"

	response, err := http.Get("https://api.vk.com/method/users.search?q=" + url.QueryEscape(name) + "&type=user&count=3&access_token=" + token)
	if err != nil {
		fmt.Printf("Error sending request to VK: %s", err)
		ch <- ChanelResult{nil, err}
	}
	defer response.Body.Close()

	result := make([]string, 3)
	if response.StatusCode == http.StatusOK {

		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Error getting answer from FB: %s", err)
			ch <- ChanelResult{nil, err}
		}
		var mapdata VKData

		err = json.Unmarshal(bodyBytes, &mapdata)
		if err != nil {
			fmt.Printf("Error unmarshaling VK answer: %s", err)
			ch <- ChanelResult{nil, err}
		}

		var person VKPerson
		for _, v := range mapdata.Response {

			err = json.Unmarshal([]byte(v), &person)
			if err != nil {
				fmt.Printf("Error unmarshaling VK answer: %s", err)
				continue
			}
			result = append(result, "https://vk.com/id"+strconv.Itoa(person.UID))
		}
	}
	ch <- ChanelResult{result, nil}
}

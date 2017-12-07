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

type VKPerson struct{
	UID int `json:"uid"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func VKSearch(name string) (result []string, err error) {
	token := "c5d5e9395af600425335104c83d9058be0eca0cca74d291d44798ca4988a358da7d833aad14058e39cf57"

	response, err := http.Get("https://api.vk.com/method/users.search?q=" + url.QueryEscape(name) + "&type=user&count=3&access_token=" + token)
	if err != nil{
		fmt.Printf("Error sending request to VK: %s", err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {

		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Error getting answer from FB: %s", err)
			return nil, err
		}
		var mapdata VKData

		err = json.Unmarshal(bodyBytes, &mapdata)
		if err != nil {
			fmt.Printf("Error unmarshaling VK answer: %s", err)
			return nil, err
		}

		var person VKPerson
		for _, v := range mapdata.Response {

			err = json.Unmarshal([]byte(v), &person)
			if err != nil{
				fmt.Printf("Error unmarshaling VK answer: %s", err)
				continue
			}
			result = append(result, "https://vk.com/id" + strconv.Itoa(person.UID))
		}
	}
	return
}
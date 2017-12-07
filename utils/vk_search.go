package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

type VKData struct {
	Response []json.RawMessage `json:"response"`
}

type VKPerson struct {
	UID       int    `json:"uid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func VKSearch(name string, commonResult *APIHandlersResult, wg *sync.WaitGroup) {
	defer wg.Done()
	token := "c5d5e9395af600425335104c83d9058be0eca0cca74d291d44798ca4988a358da7d833aad14058e39cf57"

	response, err := http.Get("https://api.vk.com/method/users.search?q=" + url.QueryEscape(name) + "&type=user&count=3&access_token=" + token)
	if err != nil {
		fmt.Printf("Error sending request to VK: %s", err)
		commonResult.Lock()
		commonResult.Errors = append(commonResult.Errors, err)
		commonResult.Unlock()
	}
	defer response.Body.Close()

	var result []string
	if response.StatusCode == http.StatusOK {

		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Error getting answer from FB: %s", err)
			commonResult.Lock()
			commonResult.Errors = append(commonResult.Errors, err)
			commonResult.Unlock()
		}
		var mapdata VKData

		err = json.Unmarshal(bodyBytes, &mapdata)
		if err != nil {
			fmt.Printf("Error unmarshaling VK answer: %s", err)
			commonResult.Lock()
			commonResult.Errors = append(commonResult.Errors, err)
			commonResult.Unlock()
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
	commonResult.Lock()
	commonResult.VK = result
	commonResult.Unlock()
}

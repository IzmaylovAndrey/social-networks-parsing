package utils

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"net/url"
	"encoding/json"
	"strconv"
)

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
		var mapdata map[string]interface{}

		err = json.Unmarshal(bodyBytes, &mapdata)
		if err != nil{
			fmt.Printf("Error unmarshaling VK answer: %s", err)
			return nil, err
		}

		for i := 1; i < len(mapdata["response"].([]interface{})); i++ {
			id := int(mapdata["response"].([]interface{})[i].(map[string]interface{})["uid"].(float64))
			result = append(result, "https://vk.com/id" + strconv.Itoa(id))
		}
	}
	return
}
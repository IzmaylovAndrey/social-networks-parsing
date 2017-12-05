package utils

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func FBSearch(name string) (result []string, err error) {
	token := "EAACEdEose0cBAHv8n40uU7KmEKCJbRMiJJ8zQT2vQlsq5q9SNcm63ZBRtYd5KJ2NZCmMYRyWZAjzm3KulmnpZAZC9XNtPPHj0WRCAZANdNRUXEk0JcvwDZCZAZAMwY6UKJxh98dVblqiug0ZCZCQsgcJlcQsH6ZAgSIgXgXsPQoZBjV9LShLFIFgXUMy4T7iLtsOCktQz0eNp7wjSWAZDZD"
	response, err := http.Get("https://graph.facebook.com/search?q=" + name + "&type=user&access_token=" + token)
	if err != nil{
		fmt.Printf("Error sending message to Telegram: %s", err)
		return nil, err
	}
	defer response.Body.Close()
	var res string
	if response.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		res = string(bodyBytes)
	}
	fmt.Printf("%s", res)
	return
}

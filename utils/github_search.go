package utils

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"net/url"
	"encoding/json"
)

type GithubData struct {
	TotalCount int `json:"total_count"`
	IncompleteResults bool `json:"incomplete_results"`
	Items []GithubPerson `json:"items"`
}

type GithubPerson struct {
	Login string `json:"login"`
	ID int `json:"id"`
	AvatarUrl string `json:"avatar_url"`
	GravatarID string `json:"gravatar_id"`
	Url string `json:"url"`
	HtmlUrl string `json:"html_url"`
	FollowersUrl string `json:"followers_url"`
	FollowingUrl string `json:"following_url"`
	GistsUrl string `json:"gists_url"`
	StarredUrl string `json:"starred_url"`
	SubscriptionsUrl string `json:"subscriptions_url"`
	OrganizationsUrl string `json:"organizations_url"`
	ReposUrl string `json:"repos_url"`
	EventsUrl string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type string `json:"type"`
	SiteAdmin bool `json:"site_admin"`
	Score float64 `json:"score"`
}

func GithubSearch(name string) (result []string, err error) {
	token := "4d721ca95335a26b2abf6a3f2b4a93093ac1a6a6"

	response, err := http.Get("https://api.github.com/search/users?q=" + url.QueryEscape(name) + "&per_page=3&access_token=" + token)
	if err != nil{
		fmt.Printf("Error sending request to Github: %s", err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		var data GithubData

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

		for _, v := range data.Items {
			result = append(result, "https://github.com/" + v.Login)
		}
	}
	return
}
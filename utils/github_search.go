package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

type GithubData struct {
	TotalCount        int            `json:"total_count"`
	IncompleteResults bool           `json:"incomplete_results"`
	Items             []GithubPerson `json:"items"`
}

type GithubPerson struct {
	Login             string  `json:"login"`
	ID                int     `json:"id"`
	AvatarUrl         string  `json:"avatar_url"`
	GravatarID        string  `json:"gravatar_id"`
	Url               string  `json:"url"`
	HtmlUrl           string  `json:"html_url"`
	FollowersUrl      string  `json:"followers_url"`
	FollowingUrl      string  `json:"following_url"`
	GistsUrl          string  `json:"gists_url"`
	StarredUrl        string  `json:"starred_url"`
	SubscriptionsUrl  string  `json:"subscriptions_url"`
	OrganizationsUrl  string  `json:"organizations_url"`
	ReposUrl          string  `json:"repos_url"`
	EventsUrl         string  `json:"events_url"`
	ReceivedEventsUrl string  `json:"received_events_url"`
	Type              string  `json:"type"`
	SiteAdmin         bool    `json:"site_admin"`
	Score             float64 `json:"score"`
}

func GithubSearch(name string, commonResult *APIHandlersResult, wg *sync.WaitGroup) {
	defer wg.Done()
	//TODO: error handling
	token := "11af0444949c31fd56e677d2f212c92976e632d7"

	response, err := http.Get("https://api.github.com/search/users?q=" + url.QueryEscape(name) + "&per_page=5&access_token=" + token)
	if err != nil {
		fmt.Printf("Error sending request to Github: %s", err)
		commonResult.Lock()
		commonResult.Errors = append(commonResult.Errors, err)
		commonResult.Unlock()
		return
	}
	defer response.Body.Close()

	var result []string
	if response.StatusCode == http.StatusOK {
		var data GithubData

		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Error getting answer from GH: %s", err)
			commonResult.Lock()
			commonResult.Errors = append(commonResult.Errors, err)
			commonResult.Unlock()
			return
		}

		err = json.Unmarshal(bodyBytes, &data)
		if err != nil {
			fmt.Printf("Error unmarshaling GH answer: %s", err)
			commonResult.Lock()
			commonResult.Errors = append(commonResult.Errors, err)
			commonResult.Unlock()
			return
		}

		for _, v := range data.Items {
			result = append(result, "https://github.com/"+v.Login)
		}
	}
	commonResult.Lock()
	commonResult.Github = result
	commonResult.Unlock()
}

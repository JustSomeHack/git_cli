package github

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/JustSomeHack/git_cli/client/http"
	"github.com/JustSomeHack/git_cli/models"
)

// GitHub interface for API client
type GitHub interface {
	GetUsers() (map[string]interface{}, error)
	GetProjects() (map[string]interface{}, error)
	PrintCommits(projectID int)
	PrintGroups()
	PrintProjects()
	PrintRequests()
	PrintUsers()
}

type gitHub struct {
	BaseURL   string
	AccessKey string
}

// NewGitHub gets a new GitHub API client
func NewGitHub(url string, key string) GitHub {
	return &gitHub{
		BaseURL:   url,
		AccessKey: key,
	}
}

func (g *gitHub) GetProjects() (map[string]interface{}, error) {
	params := &models.HTTPParams{
		URL:                 g.BaseURL + "/user/memberships/orgs",
		AuthorizationBearer: g.AccessKey,
		Headers: map[string]string{
			"Accept": "application/vnd.github.v3+json",
		},
	}

	httpClient := http.NewHTTPClient(params)
	res, err := httpClient.Get()
	if err != nil {
		return nil, err
	}

	fmt.Printf("%s\n", res)
	users := make(map[string]interface{}, 0)
	err = json.Unmarshal(res, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (g *gitHub) GetUsers() (map[string]interface{}, error) {
	params := &models.HTTPParams{
		URL:                 g.BaseURL + "/user/memberships/orgs",
		AuthorizationBearer: g.AccessKey,
		Headers: map[string]string{
			"Accept": "application/vnd.github.v3+json",
		},
	}

	httpClient := http.NewHTTPClient(params)
	res, err := httpClient.Get()
	if err != nil {
		return nil, err
	}

	fmt.Printf("%s\n", res)
	users := make(map[string]interface{}, 0)
	err = json.Unmarshal(res, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (g *gitHub) PrintCommits(projectID int) {

}

func (g *gitHub) PrintGroups() {

}

func (g *gitHub) PrintProjects() {

}

func (g *gitHub) PrintRequests() {

}

func (g *gitHub) PrintUsers() {
	users, err := g.GetUsers()
	if err != nil {
		log.Fatal(err)
	}
	for key, value := range users {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
}

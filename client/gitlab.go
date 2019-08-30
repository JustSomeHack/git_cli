package client

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/JustSomeHack/git_cli/models"
	"github.com/JustSomeHack/git_cli/models/gitlab"
)

// GitLab interface for client
type GitLab interface {
	CreateMergeRequest(target *string, source *string, title string, description string)
	GetMergeRequests(projectID int) ([]gitlab.MergeRequest, error)
	GetProjects() ([]gitlab.Project, error)
	GetUsers() ([]gitlab.User, error)
}

type gitLab struct {
	BaseURL   string
	AccessKey string
}

// NewGitLab gets a new GitLab API client
func NewGitLab(url string, key string) GitLab {
	return &gitLab{
		BaseURL:   url + "/api/v4",
		AccessKey: key,
	}
}

func (g *gitLab) CreateMergeRequest(target *string, source *string, title string, description string) {

}

func (g *gitLab) GetMergeRequests(projectID int) ([]gitlab.MergeRequest, error) {
	mergeRequestURL := fmt.Sprintf("%s/projects/%d/merge_requests", g.BaseURL, projectID)
	if projectID == -1 {
		mergeRequestURL = fmt.Sprintf("%s/merge_requests", g.BaseURL)
	}
	params := &models.HTTPParams{
		URL:                 mergeRequestURL,
		AuthorizationBearer: g.AccessKey,
	}

	httpClient := NewHTTPClient(params)
	res, err := httpClient.Get()
	if err != nil {
		log.Fatal(err)
	}
	requests := make([]gitlab.MergeRequest, 0)

	err = json.Unmarshal(res, &requests)
	if err != nil {
		log.Fatal(err)
	}
	return requests, nil
}

func (g *gitLab) GetProjects() ([]gitlab.Project, error) {
	params := &models.HTTPParams{
		URL:                 g.BaseURL + "/projects",
		AuthorizationBearer: g.AccessKey,
	}

	httpClient := NewHTTPClient(params)
	res, err := httpClient.Get()
	if err != nil {
		log.Fatal(err)
	}
	projects := make([]gitlab.Project, 0)

	err = json.Unmarshal(res, &projects)
	if err != nil {
		log.Fatal(err)
	}
	return projects, nil
}

func (g *gitLab) GetUsers() ([]gitlab.User, error) {
	params := &models.HTTPParams{
		URL:                 g.BaseURL + "/users",
		AuthorizationBearer: g.AccessKey,
	}

	httpClient := NewHTTPClient(params)
	res, err := httpClient.Get()
	if err != nil {
		log.Fatal(err)
	}
	users := make([]gitlab.User, 0)

	err = json.Unmarshal(res, &users)
	if err != nil {
		log.Fatal(err)
	}
	return users, nil
}

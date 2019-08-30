package client

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/JustSomeHack/git_cli/models"
	"github.com/JustSomeHack/git_cli/models/gitlab"
	"github.com/ryanuber/columnize"
)

// GitLab interface for client
type GitLab interface {
	CreateMergeRequest(target *string, source *string, title string, description string)
	GetMergeRequests(projectID int) ([]gitlab.MergeRequest, error)
	GetProjects() ([]gitlab.Project, error)
	GetUsers() ([]gitlab.User, error)
	PrintMergeRequests()
	PrintProjects()
	PrintUsers()
	getProject(projects []gitlab.Project, projectID int) *gitlab.Project
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

func (g *gitLab) PrintMergeRequests() {
	projects, err := g.GetProjects()
	if err != nil {
		log.Fatal(err)
	}

	requests, err := g.GetMergeRequests(-1)
	if err != nil {
		log.Fatal(err)
	}
	requestsOutput := make([]string, 0)
	requestsOutput = append(requestsOutput, "Request ID | Target Project | Source Project | Title | Description")
	for _, request := range requests {
		source := g.getProject(projects, request.SourceProjectID)
		target := g.getProject(projects, request.TargetProjectID)
		requestLine := fmt.Sprintf("%d | %s | %s | %s | %s", request.IID, target.PathWithNamespace, source.PathWithNamespace, request.Title, request.Description)
		requestsOutput = append(requestsOutput, requestLine)
	}
	fmt.Println(columnize.SimpleFormat(requestsOutput))
}

func (g *gitLab) PrintProjects() {
	projects, err := g.GetProjects()
	if err != nil {
		log.Fatal(err)
	}
	projectsOutput := make([]string, 0)
	projectsOutput = append(projectsOutput, "Project ID | Path | Description")
	for _, project := range projects {
		projectLine := fmt.Sprintf("%d | %s | %s", project.ID, project.PathWithNamespace, project.Description)
		projectsOutput = append(projectsOutput, projectLine)
	}
	fmt.Println(columnize.SimpleFormat(projectsOutput))
}

func (g *gitLab) PrintUsers() {
	users, err := g.GetUsers()
	if err != nil {
		log.Fatal(err)
	}
	usersOutput := make([]string, 0)
	usersOutput = append(usersOutput, "User ID | Name | Username | Status")
	for _, user := range users {
		if user.UserName == "ghost" || user.UserName == "root" {
			continue
		}
		userLine := fmt.Sprintf("%d | %s | %s | %s", user.ID, user.Name, user.UserName, user.State)
		usersOutput = append(usersOutput, userLine)
	}
	fmt.Println(columnize.SimpleFormat(usersOutput))
}

func (g *gitLab) getProject(projects []gitlab.Project, projectID int) *gitlab.Project {
	for _, project := range projects {
		if project.ID == projectID {
			return &project
		}
	}
	return nil
}

package gitlab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/JustSomeHack/git_cli/client/http"
	"github.com/JustSomeHack/git_cli/models"
	"github.com/JustSomeHack/git_cli/models/gitlab"
	"github.com/ryanuber/columnize"
)

// GitLab interface for http
type GitLab interface {
	AcceptRequest(projectID int, requestID int, message string, deleteSource bool)
	CreateRequest(projectID int, targetProjectID int, assigneeID int, source string, target string, title string, description string)
	GetCommits(projectID int) ([]gitlab.Commit, error)
	GetGroups() ([]gitlab.Group, error)
	GetProjects() ([]gitlab.Project, error)
	GetRequests(projectID int) ([]gitlab.Request, error)
	GetUsers() ([]gitlab.User, error)
	PrintCommits(projectID int)
	PrintGroups()
	PrintProjects()
	PrintRequests()
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

func (g *gitLab) AcceptRequest(projectID int, requestID int, message string, deleteSource bool) {
	acceptURL := fmt.Sprintf("%s/projects/%d/merge_requests/%d/merge", g.BaseURL, projectID, requestID)
	params := &models.HTTPParams{
		URL:                 acceptURL,
		AuthorizationBearer: g.AccessKey,
	}

	request := &gitlab.AcceptRequest{
		ID:              projectID,
		MergeRequestIID: requestID,
	}
	if message != "" {
		request.MergeCommitMessage = message
	}
	if deleteSource {
		request.ShouldRemoveSourceBranch = deleteSource
	}

	requestBytes, _ := json.Marshal(request)
	httpClient := http.NewHTTPClient(params)
	res, err := httpClient.Put(bytes.NewReader(requestBytes))
	if err != nil {
		log.Fatal(err)
	}
	response := new(gitlab.Request)
	err = json.Unmarshal(res, response)
	fmt.Printf("\nMerged request %d from %s into %s\n\n", requestID, response.SourceBranch, response.TargetBranch)
}

func (g *gitLab) CreateRequest(projectID int, targetProjectID int, assigneeID int, source string, target string, title string, description string) {
	createURL := fmt.Sprintf("%s/projects/%d/merge_requests", g.BaseURL, projectID)
	params := &models.HTTPParams{
		URL:                 createURL,
		AuthorizationBearer: g.AccessKey,
	}

	request := &gitlab.MergeRequest{
		ID:                 projectID,
		SourceBranch:       source,
		TargetBranch:       target,
		Title:              title,
		RemoveSourceBranch: true,
	}
	if targetProjectID > 0 {
		request.TargetProjectID = targetProjectID
	}

	if assigneeID > 0 {
		request.Assignee = assigneeID
	}
	if description != "" {
		request.Description = description
	}

	requestBytes, _ := json.Marshal(request)
	httpClient := http.NewHTTPClient(params)
	res, err := httpClient.Post(bytes.NewReader(requestBytes))
	if err != nil {
		log.Fatal(err)
	}
	response := new(gitlab.Request)
	err = json.Unmarshal(res, response)
	fmt.Printf("\nCreated request %d at %s\n\n", response.IID, response.WebURL)
}

func (g *gitLab) GetCommits(projectID int) ([]gitlab.Commit, error) {
	commitsURL := fmt.Sprintf("%s/projects/%d/repository/commits?with_stats=true", g.BaseURL, projectID)
	params := &models.HTTPParams{
		URL:                 commitsURL,
		AuthorizationBearer: g.AccessKey,
	}

	httpClient := http.NewHTTPClient(params)
	res, err := httpClient.Get()
	if err != nil {
		return nil, err
	}

	commits := make([]gitlab.Commit, 0)
	err = json.Unmarshal(res, &commits)
	if err != nil {
		return nil, err
	}
	return commits, nil
}

func (g *gitLab) GetGroups() ([]gitlab.Group, error) {
	params := &models.HTTPParams{
		URL:                 g.BaseURL + "/groups",
		AuthorizationBearer: g.AccessKey,
	}

	httpClient := http.NewHTTPClient(params)
	res, err := httpClient.Get()
	if err != nil {
		return nil, err
	}

	groups := make([]gitlab.Group, 0)
	err = json.Unmarshal(res, &groups)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (g *gitLab) GetProjects() ([]gitlab.Project, error) {
	params := &models.HTTPParams{
		URL:                 g.BaseURL + "/projects",
		AuthorizationBearer: g.AccessKey,
	}

	httpClient := http.NewHTTPClient(params)
	res, err := httpClient.Get()
	if err != nil {
		return nil, err
	}

	projects := make([]gitlab.Project, 0)
	err = json.Unmarshal(res, &projects)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (g *gitLab) GetRequests(projectID int) ([]gitlab.Request, error) {
	mergeRequestURL := fmt.Sprintf("%s/projects/%d/merge_requests?scope=all", g.BaseURL, projectID)
	if projectID == -1 {
		mergeRequestURL = fmt.Sprintf("%s/merge_requests?scope=all", g.BaseURL)
	}
	params := &models.HTTPParams{
		URL:                 mergeRequestURL,
		AuthorizationBearer: g.AccessKey,
	}

	httpClient := http.NewHTTPClient(params)
	res, err := httpClient.Get()
	if err != nil {
		return nil, err
	}

	requests := make([]gitlab.Request, 0)
	err = json.Unmarshal(res, &requests)
	if err != nil {
		return nil, err
	}
	return requests, nil
}

func (g *gitLab) GetUsers() ([]gitlab.User, error) {
	params := &models.HTTPParams{
		URL:                 g.BaseURL + "/users",
		AuthorizationBearer: g.AccessKey,
	}

	httpClient := http.NewHTTPClient(params)
	res, err := httpClient.Get()
	if err != nil {
		return nil, err
	}

	users := make([]gitlab.User, 0)
	err = json.Unmarshal(res, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (g *gitLab) PrintCommits(projectID int) {
	commits, err := g.GetCommits(projectID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("\n")
	commitsOutput := make([]string, 0)
	commitsOutput = append(commitsOutput, "Commit ID | Title | Message | Committer | Committer Email | Additions | Deletions | Date")
	for _, commit := range commits {
		commitLine := fmt.Sprintf("%s | %s | %s | %s | %s | %d | %d | %s", commit.ShortID, commit.Title, commit.Message, commit.CommitterName, commit.CommitterEmail, commit.Stats.Additions, commit.Stats.Deletions, commit.CommittedDate)
		commitsOutput = append(commitsOutput, commitLine)
	}
	fmt.Println(columnize.SimpleFormat(commitsOutput))
	fmt.Print("\n\n")
}

func (g *gitLab) PrintGroups() {
	groups, err := g.GetGroups()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("\n")
	groupsOutput := make([]string, 0)
	groupsOutput = append(groupsOutput, "Group ID | Name | Full Name | Web URL | Description")
	for _, group := range groups {
		groupLine := fmt.Sprintf("%d | %s | %s | %s | %s", group.ID, group.Name, group.FullName, group.WebURL, group.Description)
		groupsOutput = append(groupsOutput, groupLine)
	}
	fmt.Println(columnize.SimpleFormat(groupsOutput))
	fmt.Print("\n\n")
}

func (g *gitLab) PrintProjects() {
	projects, err := g.GetProjects()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("\n")
	projectsOutput := make([]string, 0)
	projectsOutput = append(projectsOutput, "Project ID | Path | SSH URL | Description")
	for _, project := range projects {
		projectLine := fmt.Sprintf("%d | %s | %s | %s", project.ID, project.PathWithNamespace, project.SSHURL, project.Description)
		projectsOutput = append(projectsOutput, projectLine)
	}
	fmt.Println(columnize.SimpleFormat(projectsOutput))
	fmt.Print("\n\n")
}

func (g *gitLab) PrintRequests() {
	projects, err := g.GetProjects()
	if err != nil {
		log.Fatal(err)
	}

	requests, err := g.GetRequests(-1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("\n")
	requestsOutput := make([]string, 0)
	requestsOutput = append(requestsOutput, "Request ID | State | Source Project | Target Project | Source Branch | Target Branch | Title | Description")
	for _, request := range requests {
		source := g.getProject(projects, request.SourceProjectID)
		target := g.getProject(projects, request.TargetProjectID)
		requestLine := fmt.Sprintf("%d | %s | %s | %s | %s | %s | %s | %s", request.IID, request.State, source.PathWithNamespace, target.PathWithNamespace, request.SourceBranch, request.TargetBranch, request.Title, request.Description)
		requestsOutput = append(requestsOutput, requestLine)
	}
	fmt.Println(columnize.SimpleFormat(requestsOutput))
	fmt.Print("\n\n")
}

func (g *gitLab) PrintUsers() {
	users, err := g.GetUsers()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("\n")
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
	fmt.Print("\n\n")
}

func (g *gitLab) getProject(projects []gitlab.Project, projectID int) *gitlab.Project {
	for _, project := range projects {
		if project.ID == projectID {
			return &project
		}
	}
	return nil
}

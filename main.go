package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/JustSomeHack/git_cli/client"
	models "github.com/JustSomeHack/git_cli/models/gitlab"
	"github.com/ryanuber/columnize"
)

var version string

func main() {
	gitURL := os.Getenv("GITHUB_URL")
	if gitURL == "" {
		gitURL = os.Getenv("GITLAB_URL")
	}

	if gitURL == "" {
		log.Fatal("GITHUB_URL or GITLAB_URL is required!")
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY is required!")
	}

	sourceProject := flag.String("source", "", "Source project [group/project]")
	targetProject := flag.String("target", "", "Target project [group/project]")

	gitlab := client.NewGitLab(gitURL, apiKey)

	users, err := gitlab.GetUsers()
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

	projects, err := gitlab.GetProjects()
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

	requests, err := gitlab.GetMergeRequests(-1)
	if err != nil {
		log.Fatal(err)
	}
	requestsOutput := make([]string, 0)
	requestsOutput = append(requestsOutput, "Request ID | Target Project | Source Project | Title | Description")
	for _, request := range requests {
		source := getProject(projects, request.SourceProjectID)
		target := getProject(projects, request.TargetProjectID)
		requestLine := fmt.Sprintf("%d | %s | %s | %s | %s", request.IID, target.PathWithNamespace, source.PathWithNamespace, request.Title, request.Description)
		requestsOutput = append(requestsOutput, requestLine)
	}
	fmt.Println(columnize.SimpleFormat(requestsOutput))

	gitlab.CreateMergeRequest(targetProject, sourceProject, "title", "description")
}

func getProject(projects []models.Project, id int) *models.Project {
	for _, project := range projects {
		if project.ID == id {
			return &project
		}
	}
	return nil
}

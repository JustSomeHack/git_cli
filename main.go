package main

import (
	"flag"
	"log"
	"os"

	"github.com/JustSomeHack/git_cli/client"
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

	gitlab.PrintUsers()

	gitlab.PrintProjects()

	gitlab.PrintMergeRequests()

	gitlab.CreateMergeRequest(targetProject, sourceProject, "title", "description")
}

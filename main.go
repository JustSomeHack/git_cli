package main

import (
	"log"
	"os"
)

var version string

func main() {
	gitLabUrl := os.Getenv("GITLAB_URL")
	gitHubUrl := os.Getenv("GITHUB_URL")

	if gitHubUrl == "" || gitLabUrl == "" {
		log.Fatal("GITHUB_URL or GITLAB_URL is required!")
	}
}

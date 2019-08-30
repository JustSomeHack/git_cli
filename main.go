package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/JustSomeHack/git_cli/client"
)

var version string

func main() {
	isGitLab := false
	gitURL := os.Getenv("GITHUB_URL")
	if gitURL == "" {
		isGitLab = true
		gitURL = os.Getenv("GITLAB_URL")
	}

	if gitURL == "" {
		log.Fatal("GITHUB_URL or GITLAB_URL is required!")
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY is required!")
	}

	acceptRequest := flag.Bool("accept-request", false, "Accept a merge/pull request")
	requestID := flag.Int("request", -1, "Merge/pull request ID, required for accept")
	message := flag.String("message", "", "Message to add to accept request")
	deleteSourceBranch := flag.Bool("delete-source", false, "Delete the source branch after accepting the request")

	createRequest := flag.Bool("create-request", false, "Create a merge/pull request")
	projectID := flag.Int("project", -1, "Project ID, required for accept and create request")
	targetProjectID := flag.Int("target-project", -1, "Target Project ID")
	assigneeID := flag.Int("assignee", -1, "Assignee ID")
	sourceBranch := flag.String("source-branch", "", "Source Branch, required for request")
	targetBranch := flag.String("target-branch", "", "Target Branch, required for request")
	title := flag.String("title", "", "Request title, required for request")
	description := flag.String("description", "", "Request description")

	listCommits := flag.Bool("list-commits", false, "List commits for a project")
	listGroups := flag.Bool("list-groups", false, "List groups")
	listProjects := flag.Bool("list-projects", false, "List projects")
	listRequests := flag.Bool("list-requests", false, "List merge or pull requests")
	listUsers := flag.Bool("list-users", false, "List users")

	flag.Parse()

	cli := client.NewCLI(isGitLab, gitURL, apiKey)

	if *acceptRequest {
		if *requestID == -1 {
			log.Fatal("Missing parameter, -request is required")
		}
		cli.AcceptRequest(*projectID, *requestID, *message, *deleteSourceBranch)
	}

	if *createRequest {
		if *sourceBranch == "" {
			log.Fatal("Missing parameter, -source-branch is required")
		}
		if *targetBranch == "" {
			log.Fatal("Missing parameter, -target-branch is required")
		}
		if *title == "" {
			log.Fatal("Missing parameter, -title is required")
		}
		cli.CreateRequest(*projectID, *targetProjectID, *assigneeID, *sourceBranch, *targetBranch, *title, *description)
	}

	if *listCommits {
		if *projectID == -1 {
			fmt.Println("Missing parameter, -project is required")
			os.Exit(10)
		}
		cli.PrintCommits(*projectID)
	}

	if *listGroups {
		cli.PrintGroups()
	}

	if *listProjects {
		cli.PrintProjects()
	}

	if *listRequests {
		cli.PrintRequests()
	}

	if *listUsers {
		cli.PrintUsers()
	}
}

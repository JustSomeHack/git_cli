package client

import (
	"github.com/JustSomeHack/git_cli/client/github"
	"github.com/JustSomeHack/git_cli/client/gitlab"
)

// CLI interface for git API
type CLI interface {
	AcceptRequest(projectID int, requestID int, message string, deleteSource bool)
	CreateRequest(projectID int, targetProjectID int, assigneeID int, source string, target string, title string, description string)
	PrintCommits(projectID int)
	PrintGroups()
	PrintProjects()
	PrintRequests()
	PrintUsers()
}

type cli struct {
	IsGitLab bool
	GitHub   github.GitHub
	GitLab   gitlab.GitLab
}

// NewCLI gets a new cli client
func NewCLI(isGitLab bool, baseURL string, accessKey string) CLI {
	return &cli{
		IsGitLab: isGitLab,
		GitHub:   github.NewGitHub(baseURL, accessKey),
		GitLab:   gitlab.NewGitLab(baseURL, accessKey),
	}
}

func (c *cli) AcceptRequest(projectID int, requestID int, message string, deleteSource bool) {
	if c.IsGitLab {
		c.GitLab.AcceptRequest(projectID, requestID, message, deleteSource)
	}
}

func (c *cli) CreateRequest(projectID int, targetProjectID int, assigneeID int, source string, target string, title string, description string) {
	if c.IsGitLab {
		c.GitLab.CreateRequest(projectID, targetProjectID, assigneeID, source, target, title, description)
	}
}

func (c *cli) PrintCommits(projectID int) {
	if c.IsGitLab {
		c.GitLab.PrintCommits(projectID)
	} else {
		c.GitHub.PrintCommits(projectID)
	}
}

func (c *cli) PrintGroups() {
	if c.IsGitLab {
		c.GitLab.PrintGroups()
	} else {
		c.GitHub.PrintGroups()
	}
}

func (c *cli) PrintProjects() {
	if c.IsGitLab {
		c.GitLab.PrintProjects()
	} else {
		c.GitHub.PrintProjects()
	}
}

func (c *cli) PrintRequests() {
	if c.IsGitLab {
		c.GitLab.PrintRequests()
	} else {
		c.GitHub.PrintRequests()
	}
}

func (c *cli) PrintUsers() {
	if c.IsGitLab {
		c.GitLab.PrintUsers()
	} else {
		c.GitHub.PrintUsers()
	}
}

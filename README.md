# git_cli

[![Build Status](https://drone.onebytedata.com/api/badges/JustSomeHack/git_cli/status.svg)](https://drone.onebytedata.com/JustSomeHack/git_cli)

Currently only works with GitLab

Must have `GITLAB_URL` and `API_KEY` in your environmental variables to work.

```bash

Usage of ./git_cli:
  -accept-request
    	Accept a merge/pull request
  -assignee int
    	Assignee ID (default -1)
  -create-request
    	Create a merge/pull request
  -delete-source
    	Delete the source branch after accepting the request
  -description string
    	Request description
  -list-commits
    	List commits for a project
  -list-groups
    	List groups
  -list-projects
    	List projects
  -list-requests
    	List merge or pull requests
  -list-users
    	List users
  -message string
    	Message to add to accept request
  -project int
    	Project ID, required for accept and create request (default -1)
  -request int
    	Merge/pull request ID, required for accept (default -1)
  -source-branch string
    	Source Branch, required for request
  -target-branch string
    	Target Branch, required for request
  -target-project int
    	Target Project ID (default -1)
  -title string
    	Request title, required for request

```

package gitlab

// CommitStat struct for GitLab response
type CommitStat struct {
	Additions int `json:"additions"`
	Deletions int `json:"deletions"`
	Total     int `json:"total"`
}

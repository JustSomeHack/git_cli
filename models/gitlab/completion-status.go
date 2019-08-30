package gitlab

// CompletionStatus struct for GitLab response
type CompletionStatus struct {
	Count          int `json:"count"`
	CompletedCount int `json:"completed_count"`
}

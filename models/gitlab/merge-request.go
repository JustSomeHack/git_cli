package gitlab

// MergeRequest struct to create merge request
type MergeRequest struct {
	ID                    int    `json:"id"`
	TargetBranch          string `json:"target_branch"`
	SourceBranch          string `json:"source_branch"`
	Title                 string `json:"title"`
	Assignee              int    `json:"assignee_id,omitempty"`
	Assignees             []int  `json:"assignee_ids,omitempty"`
	Description           string `json:"description,omitempty"`
	TargetProjectID       int    `json:"target_project_id,omitempty"`
	Labels                string `json:"labels,omitempty"`
	MilestoneID           int    `json:"milestone_id,omitempty"`
	RemoveSourceBranch    bool   `json:"remove_source_branch,omitempty"`
	AllowCollaboration    bool   `json:"allow_collaboration,omitempty"`
	AllowMaintainerToPush bool   `json:"allow_maintainer_to_push,omitempty"`
	Squash                bool   `json:"squash,omitempty"`
}

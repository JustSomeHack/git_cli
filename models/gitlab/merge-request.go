package gitlab

import "time"

// MergeRequest struct for GitLab response
type MergeRequest struct {
	ID                        int              `json:"id"`
	IID                       int              `json:"iid"`
	ProjectID                 int              `json:"project_id"`
	Title                     string           `json:"title"`
	Description               string           `json:"description"`
	State                     string           `json:"state"`
	MergedBy                  User             `json:"merged_by"`
	MergedAt                  time.Time        `json:"merged_at"`
	ClosedBy                  User             `json:"closed_by"`
	ClosedAt                  time.Time        `json:"closed_at"`
	CreatedAt                 time.Time        `json:"created_at"`
	UpdatedAt                 time.Time        `json:"updated_at"`
	TargetBranch              string           `json:"target_branch"`
	SourceBranch              string           `json:"source_branch"`
	UpVotes                   int              `json:"upvotes"`
	DownVotes                 int              `json:"downvotes"`
	Author                    User             `json:"author"`
	Assignee                  User             `json:"assignee"`
	Assignees                 []User           `json:"assignees"`
	SourceProjectID           int              `json:"source_project_id"`
	TargetProjectID           int              `json:"target_project_id"`
	Labels                    []string         `json:"labels"`
	WorkInProgress            bool             `json:"work_in_progress"`
	Milestone                 Milestone        `json:"milestone"`
	MergeWhenPipelineSucceeds bool             `json:"merge_when_pipeline_succeeds"`
	MergeStatus               string           `json:"merge_status"`
	SHA                       string           `json:"sha"`
	MergeCommitSHA            string           `json:"merge_commit_sha"`
	UserNotesCount            int              `json:"user_notes_count"`
	DiscussionLocked          bool             `json:"discussion_locked"`
	ShouldRemoveSourceBranch  bool             `json:"should_remove_source_branch"`
	ForceRemoveSourceBranch   bool             `json:"force_remove_source_branch"`
	AllowCollaboration        bool             `json:"allow_collaboration"`
	AllowMaintainerToPush     bool             `json:"allow_maintainer_to_push"`
	WebURL                    string           `json:"web_url"`
	TimeStats                 TimeStat         `json:"time_stats"`
	Squash                    bool             `json:"squash"`
	TaskCompletionStatus      CompletionStatus `json:"task_completion_status"`
}

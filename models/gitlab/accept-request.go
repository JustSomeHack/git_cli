package gitlab

// AcceptRequest struct to accept a merge request
type AcceptRequest struct {
	ID                        int    `json:"id"`
	MergeRequestIID           int    `json:"merge_request_iid"`
	MergeCommitMessage        string `json:"merge_commit_message,omitempty"`
	SquashCommitMessage       string `json:"squash_commit_message,omitempty"`
	Squash                    bool   `json:"squash,omitempty"`
	ShouldRemoveSourceBranch  bool   `json:"should_remove_source_branch,omitempty"`
	MergeWhenPipelineSucceeds bool   `json:"merge_when_pipeline_succeeds,omitempty"`
	SHA                       string `json:"sha,omitempty"`
}

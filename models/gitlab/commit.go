package gitlab

import "time"

// Commit struct for GitLab response
type Commit struct {
	ID             string     `json:"id"`
	ShortID        string     `json:"short_id"`
	Title          string     `json:"title"`
	AuthorName     string     `json:"author_name"`
	AuthorEmail    string     `json:"author_email"`
	AuthoredDate   time.Time  `json:"authored_date"`
	CommitterName  string     `json:"committer_name"`
	CommitterEmail string     `json:"committer_email"`
	CommittedDate  time.Time  `json:"committed_date"`
	CreatedAt      time.Time  `json:"created_at"`
	Message        string     `json:"message"`
	ParentIDs      []string   `json:"parent_ids"`
	Stats          CommitStat `json:"stats"`
}

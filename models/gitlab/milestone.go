package gitlab

import "time"

// Milestone GitLab milestone response
type Milestone struct {
	ID          int       `json:"id"`
	IID         int       `json:"iid"`
	ProjectID   int       `json:"project_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	State       string    `json:"state"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DueDate     time.Time `json:"due_date"`
	StartDate   time.Time `json:"start_date"`
	WebURL      string    `json:"web_url"`
}

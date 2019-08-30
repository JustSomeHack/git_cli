package github

import "time"

// Project struct for GitHub response
type Project struct {
	ID         int       `json:"id"`
	OwnerURL   string    `json:"owner_url"`
	URL        string    `json:"url"`
	HTMLURL    string    `json:"html_url"`
	ColumnsURL string    `json:"columns_url"`
	NodeID     string    `json:"node_id"`
	Name       string    `json:"name"`
	Body       string    `json:"body"`
	Number     int       `json:"number"`
	State      string    `json:"state"`
	Creator    Creator   `json:"creator"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

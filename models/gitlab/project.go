package gitlab

import "time"

// Project struct for GitLab Response
type Project struct {
	ID                int       `json:"id"`
	Description       string    `json:"description"`
	DefaultBranch     string    `json:"default_branch"`
	SSHURL            string    `json:"ssh_url_to_repo"`
	HTTPURL           string    `json:"http_url_to_repo"`
	WebURL            string    `json:"web_url"`
	ReadmeURL         string    `json:"readme_url"`
	TagList           []string  `json:"tag_list"`
	Name              string    `json:"name"`
	NameWithNamespace string    `json:"name_with_namespace"`
	Path              string    `json:"path"`
	PathWithNamespace string    `json:"path_with_namespace"`
	CreatedAt         time.Time `json:"created_at"`
	LastActivityAt    time.Time `json:"last_activity_at"`
	ForksCount        int       `json:"forks_count"`
	AvatarURL         string    `json:"avatar_url"`
	StarCount         int       `json:"star_count"`
}

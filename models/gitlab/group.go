package gitlab

// Group struct for GitLab response
type Group struct {
	ID                    int    `json:"id"`
	Name                  string `json:"name"`
	Path                  string `json:"path"`
	Description           string `json:"description"`
	Visibility            string `json:"visibility"`
	LFSEnabled            bool   `json:"lfs_enabled"`
	AvatarURL             string `json:"avatar_url"`
	WebURL                string `json:"web_url"`
	RequestAccessEnabled  bool   `json:"request_access_enabled"`
	FullName              string `json:"full_name"`
	FullPath              string `json:"full_path"`
	FileTemplateProjectID int    `json:"file_template_project_id"`
	ParentID              int    `json:"parent_id"`
}

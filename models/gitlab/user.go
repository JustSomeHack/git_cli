package gitlab

// User GitLab user struct
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	UserName  string `json:"username"`
	State     string `json:"state"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}

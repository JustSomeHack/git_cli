package gitlab

// TimeStat struct for GitLab response
type TimeStat struct {
	TimeEstimate        int `json:"time_estimate"`
	TotalTimeSpent      int `json:"total_time_spent"`
	HumanTimeEstimate   int `json:"human_time_estimate"`
	HumanTotalTimeSpent int `json:"human_total_time_spent"`
}

package model

type Applicant struct {
	Name             string   `json:"name"`
	Age              int      `json:"age"`
	Gender           string   `json:"gender"`
	HealthConditions []string `json:"health_conditions"`
	Location         string   `json:"location"`
}

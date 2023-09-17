package config

type Applicant struct {
	MinAge         int
	ValidGenders   []string
	ValidLocations []string
}

func NewApplicant() *Applicant {
	return &Applicant{
		MinAge:         18,
		ValidGenders:   []string{"m"},
		ValidLocations: []string{"London"},
	}
}

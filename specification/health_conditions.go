package specification

import "github.com/CamTrew/specification-pattern-example/model"

type HealthConditions struct{}

func (h HealthConditions) IsSatisfiedBy(a model.Applicant) bool {
	return len(a.HealthConditions) == 0
}

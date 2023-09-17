package specification

import (
	"slices"

	"github.com/thegodev/specification-pattern-example/model"
)

type Gender struct {
	validGenders []string
}

func NewGender(validGenders []string) Gender {
	return Gender{
		validGenders: validGenders,
	}
}

func (g Gender) IsSatisfiedBy(a model.Applicant) bool {
	return slices.Contains(g.validGenders, a.Gender)
}

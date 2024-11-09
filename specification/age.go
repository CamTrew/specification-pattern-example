package specification

import "github.com/CamTrew/specification-pattern-example/model"

type Age struct {
	minAge int
}

func NewAge(minAge int) Age {
	return Age{
		minAge: minAge,
	}
}

func (s Age) IsSatisfiedBy(a model.Applicant) bool {
	return a.Age >= s.minAge
}

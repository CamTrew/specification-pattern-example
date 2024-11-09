package specification

import "github.com/CamTrew/specification-pattern-example/model"

type Applicant interface {
	IsSatisfiedBy(a model.Applicant) bool
}

type AndApplicant struct {
	specs []Applicant
}

func (a AndApplicant) IsSatisfiedBy(app model.Applicant) bool {
	for _, spec := range a.specs {
		if !spec.IsSatisfiedBy(app) {
			return false
		}
	}
	return true
}

func NewAndApplicant(specs ...Applicant) AndApplicant {
	return AndApplicant{
		specs: specs,
	}
}

package specification

import "github.com/thegodev/specification-pattern-example/model"

type Applicant interface {
	IsSatisfiedBy(a model.Applicant) bool
}

type OrApplicant struct {
	specs []Applicant
}

func (o OrApplicant) IsSatisfiedBy(a model.Applicant) bool {
	for _, spec := range o.specs {
		if !spec.IsSatisfiedBy(a) {
			return false
		}
	}
	return true
}

func NewOrApplicant(specs ...Applicant) OrApplicant {
	return OrApplicant{
		specs: specs,
	}
}

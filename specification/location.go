package specification

import (
	"slices"

	"github.com/CamTrew/specification-pattern-example/model"
)

type Location struct {
	validLocations []string
}

func NewLocation(validLocations []string) Location {
	return Location{
		validLocations: validLocations,
	}
}

func (l Location) IsSatisfiedBy(a model.Applicant) bool {
	return slices.Contains(l.validLocations, a.Location)
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/CamTrew/specification-pattern-example/config"
	"github.com/CamTrew/specification-pattern-example/model"
	"github.com/CamTrew/specification-pattern-example/specification"
)

func filter(as []model.Applicant, spec specification.Applicant) []model.Applicant {
	result := make([]model.Applicant, 0, len(as))
	for _, a := range as {
		if spec.IsSatisfiedBy(a) {
			result = append(result, a)
		}
	}
	return result
}

func main() {
	f, err := os.ReadFile("./applicants.json")
	if err != nil {
		log.Fatal("unable to read file")
	}

	var applicants []model.Applicant
	if err := json.Unmarshal(f, &applicants); err != nil {
		log.Fatal("unable to marshal into struct")
	}

	cfg := config.NewApplicant()

	spec := specification.NewAndApplicant(
		specification.NewAge(cfg.MinAge),
		specification.NewGender(cfg.ValidGenders),
		specification.NewLocation(cfg.ValidLocations),
		specification.HealthConditions{},
	)

	filtered := filter(applicants, spec)

	for _, applicant := range filtered {
		fmt.Println(applicant.Name)
	}
}

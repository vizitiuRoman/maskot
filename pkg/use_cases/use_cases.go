package use_cases

import (
	"github.com/maskot/pkg/repositories"
	"github.com/maskot/pkg/use_cases/seamless"
)

type UseCases struct {
	Seamless *seamless.UseCases
}

type Dependencies struct {
	Repos *repositories.Repository
}

func NewUseCases(deps *Dependencies) *UseCases {
	return &UseCases{
		Seamless: seamless.NewUseCases(deps.Repos),
	}
}

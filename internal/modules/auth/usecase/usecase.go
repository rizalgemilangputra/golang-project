package usecase

import "golang-project/internal/modules/auth/repository"

type Usecase struct {
	Repository *repository.Repository
}

func NewUsecase(repository *repository.Repository) *Usecase {
	return &Usecase{
		Repository: repository,
	}
}

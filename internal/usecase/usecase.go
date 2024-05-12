package usecase

import "github.com/MamushevArup/ds-voice/internal/repository"

type UseCase struct {
	repo *repository.Repository
}

func New() *UseCase {
	return &UseCase{}
}

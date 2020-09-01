package usecase

import "cleanarchitecture-fes/src/domain"

type FesEventRepository interface {
	Create(domain.FesEvent) error
	GetAll() (*domain.FesEvents, error)
}

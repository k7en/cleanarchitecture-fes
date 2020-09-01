package feseventinteractor

import (
	"cleanarchitecture-fes/src/domain"
	"cleanarchitecture-fes/src/usecase"
)

type FesEeventInteractor struct {
	fesEventRepository usecase.FesEventRepository
}

func New(
	fesEventRepository usecase.FesEventRepository,
) *FesEeventInteractor {
	return &FesEeventInteractor{
		fesEventRepository: fesEventRepository,
	}
}

func (interactor *FesEeventInteractor) Save(fesEvent domain.FesEvent) error {
	err := interactor.fesEventRepository.Create(fesEvent)
	if err != nil {
		return err
	}
	return nil
}

func (interactor *FesEeventInteractor) Get() (*domain.FesEvents, error) {
	createdFesEvent, err := interactor.fesEventRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return createdFesEvent, nil
}

package feseventinteractor_test

import (
	"cleanarchitecture-fes/src/domain"
	"cleanarchitecture-fes/src/usecase/feseventinteractor"
	"errors"
	"log"
	"testing"

	"github.com/tj/assert"
)

func TestFesEventSaveUsecase(t *testing.T) {
	fesEventRepository := NewFesEventRepositoryMock()
	uc := feseventinteractor.New(fesEventRepository)
	inputFesEvent := domain.FesEvent{
		Title:   "FesEventTestDummy",
		Speaker: "FesEventSpeakerDummy",
	}
	fesEventRepository.On(
		"Create",
		inputFesEvent,
	).Return(nil).Once()
	err := uc.Save(inputFesEvent)
	assert.Nil(t, err)
}
func TestFesEventSaveErrorUsecase(t *testing.T) {
	fesEventRepository := NewFesEventRepositoryMock()
	uc := feseventinteractor.New(fesEventRepository)
	inputFesEvent := domain.FesEvent{
		Title:   "FesEventTestDummyError",
		Speaker: "FesEventSpeakerDummyError",
	}
	fesEventRepository.On(
		"Create",
		inputFesEvent,
	).Return(errors.New("RepositoryError!")).Once()
	err := uc.Save(inputFesEvent)
	assert.Error(t, err)
}
func TestFesEventGetUsecase(t *testing.T) {
	fesEventRepository := NewFesEventRepositoryMock()
	uc := feseventinteractor.New(fesEventRepository)
	retFesEvents := domain.FesEvents{
		{
			ID:      1,
			Title:   "TestTitle01",
			Speaker: "Speaker01",
		},
		{
			ID:      2,
			Title:   "TestTitle02",
			Speaker: "Speaker02",
		},
		{
			ID:      3,
			Title:   "TestTitle03",
			Speaker: "Speaker03",
		},
	}
	fesEventRepository.On(
		"GetAll",
	).Return(&retFesEvents, nil).Once()
	fesEvents, err := uc.Get()
	assert.Nil(t, err)
	log.Printf("GetFesEvents: %v", fesEvents)
}

func TestFesEventGetErrorUsecase(t *testing.T) {
	fesEventRepository := NewFesEventRepositoryMock()
	uc := feseventinteractor.New(fesEventRepository)
	fesEventRepository.On(
		"GetAll",
	).Return(nil, errors.New("GetAll Error!")).Once()
	fesEvents, err := uc.Get()
	assert.Nil(t, fesEvents)
	assert.Error(t, err)
}

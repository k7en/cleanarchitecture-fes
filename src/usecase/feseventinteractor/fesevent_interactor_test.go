package feseventinteractor_test

import (
	"cleanarchitecture-fes/src/adaptor/repository"
	"cleanarchitecture-fes/src/domain"
	"cleanarchitecture-fes/src/usecase/feseventinteractor"
	"log"
	"os"
	"testing"

	"github.com/tj/assert"
)

func TestFesEventSaveUsecase(t *testing.T) {
	args := os.Getenv("MYSQL_ARGS")
	if args == "" {
		panic("Please set MYSQL_ARGS env")
	}
	fesEventRepository, err := repository.New(args)
	if err != nil {
		panic("repository init error")
	}
	uc := feseventinteractor.New(fesEventRepository)
	inputFesEvent := domain.FesEvent{
		Title:   "FesEventTestDummy",
		Speaker: "FesEventSpeakerDummy",
	}
	err = uc.Save(inputFesEvent)
	assert.Nil(t, err)
}

func TestFesEventGetUsecase(t *testing.T) {
	args := os.Getenv("MYSQL_ARGS")
	if args == "" {
		panic("Please set MYSQL_ARGS env")
	}
	fesEventRepository, err := repository.New(args)
	if err != nil {
		panic("repository init error")
	}
	uc := feseventinteractor.New(fesEventRepository)
	fesEvents, err := uc.Get()
	assert.Nil(t, err)
	log.Printf("GetFesEvents: %v", fesEvents)
}

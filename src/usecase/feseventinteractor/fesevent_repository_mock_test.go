package feseventinteractor_test

// import (
// 	"cleanarchitecture-fes/src/domain"

// 	"github.com/stretchr/testify/mock"
// )

// type FesEventRepositoryMock struct {
// 	mock.Mock
// }

// func NewFesEventRepositoryMock() *FesEventRepositoryMock {
// 	return &FesEventRepositoryMock{}
// }

// func (r *FesEventRepositoryMock) Create(domain.FesEvent) (domain.FesEvent, error) {
// 	ret := r.Called()
// 	return ret.Get(0).(domain.FesEvent), ret.Error(1)
// }

// func (r *FesEventRepositoryMock) GetAll() (domain.FesEvents, error) {
// 	ret := r.Called()
// 	return ret.Get(0).(domain.FesEvents), ret.Error(1)
// }
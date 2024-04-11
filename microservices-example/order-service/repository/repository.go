package repository

import (
	"sync"

	"my.go/microservicesgrpc/service/order/model"
)

var (
	repo *MockRepository
)

type MockRepository struct {
	Data []model.Order
	m    sync.Mutex
}

func (r *MockRepository) GetAll() []model.Order {
	return r.Data
}

func (r *MockRepository) Insert(order model.Order) {
	r.m.Lock()
	defer r.m.Unlock()

	r.Data = append(r.Data, order)
}

func MockRepositoryInstance() *MockRepository {
	if repo == nil {
		repo = &MockRepository{
			Data: make([]model.Order, 0),
		}
	}
	return repo
}

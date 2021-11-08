package memsql

import (
	"helloworld/domain"
)

type memSqlUserRepository struct {
	dbString string
}

func NewMemSqlUserRepository(dbStr string) domain.UserRepository {
	return &memSqlUserRepository{dbStr}
}

func (memUr *memSqlUserRepository) GetID(id string) (error, domain.User){
	d := domain.User{ID: "memsql testing ID", Name: "memsql testing name"}

	return nil, d
}

func (memUr *memSqlUserRepository) GetName(name string) (error, domain.User) {
	d := domain.User{Name: "memsql testing name", ID: "memsql testing ID"}

	return nil, d
}
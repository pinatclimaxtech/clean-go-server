package mysql

import (
	"helloworld/domain"
)

type mysqlUserRepository struct {
	dbString string
}

func NewMySQLUserRepository(dbStr string) domain.UserRepository {
	return &mysqlUserRepository{dbStr}
}

func (mur *mysqlUserRepository) GetID(id string) (error, domain.User){
	d := domain.User{ID: "mysql testing ID", Name: "mysql testing name"}

	return nil, d
}

func (mur *mysqlUserRepository) GetName(name string) (error, domain.User) {
	d := domain.User{Name: "mysql testing name", ID: "mysql testing ID"}

	return nil, d
}
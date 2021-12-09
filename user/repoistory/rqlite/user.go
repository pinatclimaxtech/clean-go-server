package rqlite

import (
	"clean-go-server/domain"
	"github.com/rqlite/gorqlite"
)

type rqliteUserRepository struct {
	sql *gorqlite.Connection
}

func NewRqliteUserRepository(sql *gorqlite.Connection) domain.UserRepository {
	return &rqliteUserRepository{sql}
}

func (mur *rqliteUserRepository) GetID(id string) (error, domain.User) {

	var uid string
	var name string
	rows, err := mur.sql.QueryOne("select uid, username from testing where uid =" + id)
	if err != nil {

	}

	for rows.Next() {
		err := rows.Scan(&uid, &name)
		if err != nil {

		}
	}
	d := domain.User{ID: uid, Name: name}

	return nil, d
}

func (mur *rqliteUserRepository) GetName(name string) (error, domain.User) {
	d := domain.User{Name: "mysql testing name", ID: "mysql testing ID"}

	return nil, d
}

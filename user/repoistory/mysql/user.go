package mysql

import (
	"clean-go-server/domain"
	"database/sql"
)

type mysqlUserRepository struct {
	sql *sql.DB
}

func NewMySQLUserRepository(sql *sql.DB) domain.UserRepository {
	return &mysqlUserRepository{sql}
}

func (mur *mysqlUserRepository) GetID(id string) (error, domain.User) {
	sqlSyntax := "SELECT uid, username FROM testing.userinfo WHERE uid=?"
	stmt, err := mur.sql.Prepare(sqlSyntax)
	defer stmt.Close()
	if err != nil {
	}

	rows, err := stmt.Query(&id)
	if err != nil {

	}

	var uid, username string
	for rows.Next() {
		if err := rows.Scan(&uid, &username); err != nil {

		}
	}
	d := domain.User{ID: uid, Name: username}

	return nil, d
}

func (mur *mysqlUserRepository) GetName(name string) (error, domain.User) {
	d := domain.User{Name: "mysql testing name", ID: "mysql testing ID"}

	return nil, d
}

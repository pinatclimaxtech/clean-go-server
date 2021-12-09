package main

import (
	"clean-go-server/user/repoistory/mysql"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Database interface {
	Connect() string
}

type MySQL struct {
	Name string
}

type MemSQL struct {
	Name string
}

func (d *MySQL) Connect() string {
	return fmt.Sprintf("%s is connecting\n", d.Name)
}

func (mem *MemSQL) Connect() string {
	return fmt.Sprintf("%s is connecting\n", mem.Name)
}

func GetConnect(database Database) string {
	return database.Connect()
}

func main() {

	user := "root"
	password := "27940001"
	ip := "192.168.50.111"
	port := "3306"

	db, err := sql.Open("mysql", user+":"+password+"@tcp("+ip+":"+port+")/?charset=utf8")
	if err != nil {
		fmt.Println("sql.Open error: ", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("ping error: ", err)
	}

	mUr := mysql.NewMySQLUserRepository(db)
	_, name := mUr.GetID("2")

	fmt.Printf("%+v\n", name)

	//memSql := MemSQL{Name: "MemSQL"}
	//s = GetConnect(&memSql)
	//memUr := memsql.NewMemSqlUserRepository(s)
	//_, name = memUr.GetName("my name")
	//
	//fmt.Printf("%+v\n", name)

}

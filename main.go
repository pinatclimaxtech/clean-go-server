package main

import (
	"fmt"
	"helloworld/user/repoistory/memsql"
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
	mysql := MySQL{Name: "mysql80"}
	s := GetConnect(&mysql)

	mUr := mysql.NewMySQLUserRepository(s)
	_, name := mUr.GetName("my name")

	fmt.Printf("%+v\n",name)

	memSql := MemSQL{Name: "MemSQL"}
	s = GetConnect(&memSql)
	memUr := memsql.NewMemSqlUserRepository(s)
	_, name = memUr.GetName("my name")

	fmt.Printf("%+v\n",name)



}

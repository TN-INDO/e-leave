package main

import (
	db "server/structs/db"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTableUsers_20180604_145259 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTableUsers_20180604_145259{}
	m.Created = "20180604_145259"

	migration.Register("CreateTableUsers_20180604_145259", m)
}

// Run the migrations
func (m *CreateTableUsers_20180604_145259) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	var table db.User
	dt := "20180604"
	m.SQL(GetQuery(table.TableName(), dt, "create_table_users"))
}

// Reverse the migrations
func (m *CreateTableUsers_20180604_145259) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	var table db.User
	dt := "20180604"
	m.SQL(GetQuery(table.TableName(), dt, "drop_table_users"))
}

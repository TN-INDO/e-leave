package main

import (
	"server/structs/db"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTableUserTypeLeave_20180704_181620 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTableUserTypeLeave_20180704_181620{}
	m.Created = "20180704_181620"

	migration.Register("CreateTableUserTypeLeave_20180704_181620", m)
}

// Run the migrations
func (m *CreateTableUserTypeLeave_20180704_181620) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	var table db.UserTypeLeave
	dt := "20180704"
	m.SQL(GetQuery(table.TableName(), dt, "create_table_user_type_leave"))
}

// Reverse the migrations
func (m *CreateTableUserTypeLeave_20180704_181620) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	var table db.UserTypeLeave
	dt := "20180704"
	m.SQL(GetQuery(table.TableName(), dt, "drop_table_user_type_leave"))
}

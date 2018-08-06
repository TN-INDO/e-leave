package main

import (
	db "server/structs/db"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTableTypeLeave_20180704_164723 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTableTypeLeave_20180704_164723{}
	m.Created = "20180704_164723"

	migration.Register("CreateTableTypeLeave_20180704_164723", m)
}

// Run the migrations
func (m *CreateTableTypeLeave_20180704_164723) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	var table db.TypeLeave
	dt := "20180704"
	m.SQL(GetQuery(table.TableName(), dt, "create_table_type_leave"))
}

// Reverse the migrations
func (m *CreateTableTypeLeave_20180704_164723) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	var table db.TypeLeave
	dt := "20180704"
	m.SQL(GetQuery(table.TableName(), dt, "drop_table_type_leave"))
}

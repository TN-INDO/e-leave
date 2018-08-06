package main

import (
	"server/structs/db"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertToUsers_20180712_103101 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertToUsers_20180712_103101{}
	m.Created = "20180712_103101"

	migration.Register("InsertToUsers_20180712_103101", m)
}

// Run the migrations
func (m *InsertToUsers_20180712_103101) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	var table db.User
	dt := "20180712"
	m.SQL(GetQuery(table.TableName(), dt, "insert_to_users"))
}

// Reverse the migrations
func (m *InsertToUsers_20180712_103101) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}

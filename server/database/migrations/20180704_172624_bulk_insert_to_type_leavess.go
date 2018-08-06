package main

import (
	"server/structs/db"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type BulkInsertToTypeLeavess_20180704_172624 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &BulkInsertToTypeLeavess_20180704_172624{}
	m.Created = "20180704_172624"

	migration.Register("BulkInsertToTypeLeavess_20180704_172624", m)
}

// Run the migrations
func (m *BulkInsertToTypeLeavess_20180704_172624) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	var table db.TypeLeave
	dt := "20180704"
	m.SQL(GetQuery(table.TableName(), dt, "bulk_insert_to_type_leave"))
}

// Reverse the migrations
func (m *BulkInsertToTypeLeavess_20180704_172624) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
}

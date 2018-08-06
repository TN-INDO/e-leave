package main

import (
	"server/structs/db"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddFkToTypeLeaveId_20180704_170807 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddFkToTypeLeaveId_20180704_170807{}
	m.Created = "20180704_170807"

	migration.Register("AddFkToTypeLeaveId_20180704_170807", m)
}

// Run the migrations
func (m *AddFkToTypeLeaveId_20180704_170807) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	var table db.TypeLeave
	dt := "20180704"
	m.SQL(GetQuery(table.TableName(), dt, "add_fk_to_type_leave_id"))
}

// Reverse the migrations
func (m *AddFkToTypeLeaveId_20180704_170807) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
}

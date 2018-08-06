package main

import (
	db "server/structs/db"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddSequenceLeaveRequest_20180629_152857 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddSequenceLeaveRequest_20180629_152857{}
	m.Created = "20180629_152857"

	migration.Register("AddSequenceLeaveRequest_20180629_152857", m)
}

// Run the migrations
func (m *AddSequenceLeaveRequest_20180629_152857) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	var table db.LeaveRequest
	dt := "20180629"
	m.SQL(GetQuery(table.TableName(), dt, "create_table_sequence_leave_request"))
}

// Reverse the migrations
func (m *AddSequenceLeaveRequest_20180629_152857) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	var table db.LeaveRequest
	dt := "20180629"
	m.SQL(GetQuery(table.TableName(), dt, "drop_table_sequence_leave_request"))
}

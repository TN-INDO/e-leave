package main

import (
	db "server/structs/db"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTableLeaveRequest_20180607_121706 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTableLeaveRequest_20180607_121706{}
	m.Created = "20180607_121706"

	migration.Register("CreateTableLeaveRequest_20180607_121706", m)
}

// Run the migrations
func (m *CreateTableLeaveRequest_20180607_121706) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	var table db.LeaveRequest
	dt := "20180607"
	m.SQL(GetQuery(table.TableName(), dt, "create_table_leave_request"))

}

// Reverse the migrations
func (m *CreateTableLeaveRequest_20180607_121706) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	var table db.LeaveRequest
	dt := "20180607"
	m.SQL(GetQuery(table.TableName(), dt, "drop_table_leave_request"))
}

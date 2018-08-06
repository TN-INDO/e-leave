package employee

import (
	dbInterfaceEmployee "server/models/db/interfaces/employee"
	dbLayerEmployee "server/models/db/pgsql/employee"
)

// constant var
var (
	DBPostEmployee dbInterfaceEmployee.IBaseEmployee
)

func init() {
	DBPostEmployee = new(dbLayerEmployee.Employee)
}

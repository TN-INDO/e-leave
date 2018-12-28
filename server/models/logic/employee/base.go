package employee

import (
	dbInterfaceEmployee "server/models/db/interfaces/employee"
	dbLayerEmployee "server/models/db/pgsql/employee"
)

// constant var
var (
	DBEmployee dbInterfaceEmployee.IBaseEmployee
)

func init() {
	DBEmployee = new(dbLayerEmployee.Employee)
}

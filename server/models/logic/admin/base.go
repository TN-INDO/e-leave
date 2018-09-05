package admin

import (
	dbInterfacAdmin "server/models/db/interfaces/admin"
	dbLayerAdmin "server/models/db/pgsql/admin"
)

// constant var
var (
	DBAdmin dbInterfacAdmin.IBaseAdmin
)

func init() {
	DBAdmin = new(dbLayerAdmin.Admin)
}

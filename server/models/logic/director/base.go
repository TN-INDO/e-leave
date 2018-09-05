package director

import (
	dbInterfaceDirector "server/models/db/interfaces/director"
	dbLayerDirector "server/models/db/pgsql/director"
)

// constant var
var (
	DBDirector dbInterfaceDirector.IBaseDirector
)

func init() {
	DBDirector = new(dbLayerDirector.Director)
}

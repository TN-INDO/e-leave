package supervisor

import (
	dbInterfaceSupervisor "server/models/db/interfaces/supervisor"
	dbLayerSupervisor "server/models/db/pgsql/supervisor"
)

// constant var
var (
	DBSupervisor dbInterfaceSupervisor.IBaseSupervisor
)

func init() {
	DBSupervisor = new(dbLayerSupervisor.Supervisor)
}

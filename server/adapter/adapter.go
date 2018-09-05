package adapter

import "server/helpers/constant"

// CallPGSQL ...
func CallPGSQL() string {
	return "user=postgres password=root host=172.17.0.1 port=5432 dbname=db_leave_request sslmode=disable"
}

// CallSQLITE ...
func CallSQLITE() string {
	return constant.GOPATH + "/src/" + constant.GOAPP + "/database/sqlite/db_leave_request.db"
}

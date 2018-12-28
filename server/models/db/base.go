package db

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	// _ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"server/adapter"
	"server/helpers"
	"server/helpers/constant"
	dbStruct "server/structs/db"
)

// RegisterPGSQL ...
func RegisterPGSQL() {
	maxIdle := 30
	maxConn := 30

	errRegisterDriver := orm.RegisterDriver("postgres", orm.DRPostgres)
	if errRegisterDriver != nil {
		helpers.CheckErr("error while register driver @RegisterPGSQL", errRegisterDriver)
	}

	errRegisterDataBase := orm.RegisterDataBase("default", "postgres",
		adapter.CallPGSQL(),
		maxIdle, maxConn)
	if errRegisterDataBase != nil {
		helpers.CheckErr("error while register DB @RegisterPGSQL", errRegisterDataBase)
	}

	RegisterModel()
}

// RegisterSQLite ...
func RegisterSQLite() {
	maxIdle := 30
	maxConn := 30

	errRegisterDriver := orm.RegisterDriver("sqlite3", orm.DRSqlite)
	if errRegisterDriver != nil {
		helpers.CheckErr("error while register driver @RegisterSQLite", errRegisterDriver)
	}

	errRegisterDataBase := orm.RegisterDataBase("default", "sqlite3", adapter.CallSQLITE(), maxIdle, maxConn)
	if errRegisterDataBase != nil {
		helpers.CheckErr("error while register DB @RegisterSQLite", errRegisterDataBase)
	}
	// beego.Debug(adapter.CallSQLITE())

	RegisterModel()
	CreateTableUser()
	CreateTableLeaveRequest()
	CreateTableTypeLeave()
	CreateTableUserTypeLeave()

	MigrateData("users")
	MigrateData("type_leave")
}

// RegisterModel to register database
func RegisterModel() {
	orm.RegisterModel(new(dbStruct.LeaveRequest))
	orm.RegisterModel(new(dbStruct.User))
	orm.RegisterModel(new(dbStruct.TypeLeave))
	orm.RegisterModel(new(dbStruct.UserTypeLeave))
}

// CreateTableUser ...
func CreateTableUser() {
	var user dbStruct.User

	o := orm.NewOrm()
	o.Using("default")

	qb := []string{
		"CREATE TABLE IF NOT EXISTS",
		user.TableName(),
		"(",
		"employee_number integer NOT NULL PRIMARY KEY,",
		"name text NOT NULL,",
		"gender text NOT NULL,",
		"position text NOT NULL,",
		"start_working_date text NOT NULL,",
		"mobile_phone text NOT NULL,",
		"email text NOT NULL,",
		"password varchar(100) NOT NULL,",
		"role text NOT NULL,",
		"supervisor_id integer,",
		"created_at timestamp with time zone NOT NULL default CURRENT_TIMESTAMP,",
		"updated_at timestamp with time zone);",
	}

	sql := strings.Join(qb, " ")
	beego.Debug(sql)

	res, err := o.Raw(sql).Exec()
	if err != nil {
		helpers.CheckErr("error creating table users @CreateTableUser", err)
	}

	beego.Debug(res)
}

// CreateTableLeaveRequest ...
func CreateTableLeaveRequest() {
	var leave dbStruct.LeaveRequest

	o := orm.NewOrm()
	o.Using("default")

	qb := []string{
		"CREATE TABLE IF NOT EXISTS",
		leave.TableName(),
		"(",
		"id integer NOT NULL PRIMARY KEY AUTOINCREMENT,",
		"employee_number integer NOT NULL,",
		"type_leave_id integer NOT NULL,",
		"reason text NOT NULL,",
		"date_from text NOT NULL,",
		"date_to text NOT NULL,",
		"half_dates text [],",
		"back_on text NOT NULL,",
		"total float NOT NULL,",
		"contact_address text NOT NULL,",
		"contact_number text NOT NULL,",
		"status text NOT NULL,",
		"action_by text,",
		"reject_reason text,",
		"created_at timestamp NOT NULL default CURRENT_TIMESTAMP,",
		"updated_at timestamp with time zone,",
		"FOREIGN KEY (employee_number) REFERENCES users (employee_number) ON DELETE CASCADE ON UPDATE CASCADE);",
	}

	sql := strings.Join(qb, " ")
	beego.Debug(sql)

	res, err := o.Raw(sql).Exec()
	if err != nil {
		helpers.CheckErr("error creating table leave_request @CreateTableLeaveRequest", err)
	}
	beego.Debug(res)
}

// CreateTableTypeLeave ...
func CreateTableTypeLeave() {
	var typeLeave dbStruct.TypeLeave

	o := orm.NewOrm()
	o.Using("default")

	qb := []string{
		"CREATE TABLE IF NOT EXISTS",
		typeLeave.TableName(),
		"(",
		"id integer PRIMARY KEY NOT NULL,",
		"type_name text NOT NULL);",
	}

	sql := strings.Join(qb, " ")
	beego.Debug(sql)

	res, err := o.Raw(sql).Exec()
	if err != nil {
		helpers.CheckErr("error creating table type_leave @CreateTableTypeLeave", err)
	}
	beego.Debug(res)
}

// CreateTableUserTypeLeave ...
func CreateTableUserTypeLeave() {
	var userTypeLeave dbStruct.UserTypeLeave

	o := orm.NewOrm()
	o.Using("default")

	qb := []string{
		"CREATE TABLE IF NOT EXISTS",
		userTypeLeave.TableName(),
		"(",
		"id integer NOT NULL PRIMARY KEY AUTOINCREMENT,",
		"employee_number integer NOT NULL,",
		"type_leave_id integer NOT NULL,",
		"leave_remaining float NOT NULL,",
		"FOREIGN KEY (employee_number) REFERENCES users (employee_number) ON DELETE CASCADE ON UPDATE CASCADE,",
		"FOREIGN KEY (type_leave_id) REFERENCES leave_request (type_leave_id) ON DELETE CASCADE ON UPDATE CASCADE);",
	}

	sql := strings.Join(qb, " ")
	beego.Debug(sql)

	res, err := o.Raw(sql).Exec()
	if err != nil {
		helpers.CheckErr("error creating table user_type_leave @CreateTableUserTypeLeave", err)
	}
	beego.Debug(res)
}

// MigrateData ...
func MigrateData(param string) {
	o := orm.NewOrm()

	if param == "type_leave" {
		var typeLeave []dbStruct.TypeLeave
		fl := constant.GOPATH + "/src/" + constant.GOAPP + "/database/sqlite/seeders/data_type_leave.json"

		raw, err := ioutil.ReadFile(fl)
		if err != nil {
			helpers.CheckErr("failed read file seeder @MigrateData", err)
		}

		err = json.Unmarshal(raw, &typeLeave)
		if err != nil {
			helpers.CheckErr("failed unmarshall seeders @MigrateData", err)
		}

		cnt, errMulti := o.InsertMulti(len(typeLeave), typeLeave)
		beego.Debug(cnt, errMulti)

	} else if param == "users" {
		var users []dbStruct.User
		fl := constant.GOPATH + "/src/" + constant.GOAPP + "/database/sqlite/seeders/data_admin.json"

		raw, err := ioutil.ReadFile(fl)
		if err != nil {
			helpers.CheckErr("failed read file seeder @MigrateData", err)
		}

		err = json.Unmarshal(raw, &users)
		if err != nil {
			helpers.CheckErr("failed unmarshall seeders @MigrateData", err)
		}

		cnt, errMulti := o.InsertMulti(len(users), users)
		beego.Debug(cnt, errMulti)
	}
}

// ResetDB ...
func ResetDB() {
	var (
		user          dbStruct.User
		leave         dbStruct.LeaveRequest
		typeLeave     dbStruct.TypeLeave
		userTypeLeave dbStruct.UserTypeLeave
	)

	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		helpers.CheckErr("Error begin", err)
	}

	res1, errRaw1 := o.Raw(`DELETE FROM ` + user.TableName()).Exec()
	if errRaw1 != nil {
		helpers.CheckErr("error reset users @ResetDB", errRaw1)
	}

	res2, errRaw2 := o.Raw(`DELETE FROM ` + leave.TableName()).Exec()
	if errRaw2 != nil {
		helpers.CheckErr("error reset leave_request @ResetDB", errRaw2)
	}

	res3, errRaw3 := o.Raw(`DELETE FROM ` + typeLeave.TableName()).Exec()
	if errRaw3 != nil {
		helpers.CheckErr("error reset type_leave @ResetDB", errRaw3)
	}

	res4, errRaw4 := o.Raw(`DELETE FROM ` + userTypeLeave.TableName()).Exec()
	if errRaw4 != nil {
		helpers.CheckErr("error reset user_type_leave @ResetDB", errRaw4)
	}

	err = o.Commit()
	if err != nil {
		helpers.CheckErr("Error commit", err)
	}

	beego.Debug(res1, res2, res3, res4)
}

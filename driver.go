package hive

import (
"database/sql"
"database/sql/driver"
)

func init(){
	sql.Register("hive", &hiveDrive{})
}

type hiveDrive struct {}

// Implements driver.Conn interface:
type hiveConn struct {
	hiveVersion string
	host        string
	port        int
	user        string
	password    string
	dbName      string
	args        map[string]interface{}
}

func (drv *hiveDrive) Open(dsn string) (driver.Conn, error) {


    return nil, nil
}

func OpenHiveconnection(dsn string) (driver.Conn, error) {
	return nil, nil
}
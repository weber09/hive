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

type hiveStatement struct {
	connection *hiveConn
	preparedQuery string
	argsCount int
}

func (drv *hiveDrive) Open(dsn string) (driver.Conn, error) {
	conn, err := OpenHiveconnection(dsn)

	if err != nil {
		return nil ,err
	}

    return conn, nil
}

func (*hiveConn) Prepare(query string) (driver.Stmt, error){
	return nil, nil
}

func (*hiveConn) Close() error {
	return nil
}

func (*hiveConn) Begin() (driver.Tx, error){
	return nil, nil
}

func OpenHiveconnection(dsn string) (driver.Conn, error) {
	config, err := ParseDsn(dsn)
	if err != nil {
		return nil, err
	}

	conn := &hiveConn{host: config.hostName, port: config.port, dbName: config.dbName,
		              args: config.args, hiveVersion: config.hiveType,
		              password: config.password, user: config.user}

	return conn, nil
}
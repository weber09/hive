package hive

import (
	"testing"
)

func TestParseDsn(t *testing.T) {
	connDsn := "jdbc:hive2://localhost:10000/default"

	conn, _ := ParseDsn(connDsn)

	if conn == nil {
		t.Errorf("Expected conn but got nil")
	}

	if len(conn.hostName) == 0 {
		t.Errorf("Expected hostName but got nil or empty")
	}

	if conn.hostName != "localhost" {
		t.Errorf("Expected host as localhost instead got %s", conn.hostName)
	}

	if conn.port == 0 {
		t.Errorf("Expected port but got empty")
	}

	if len(conn.dbName) == 0 {
		t.Errorf("Expected dbName but got empty")
	}

	if conn.dbName != "default" {
		t.Errorf("Expected dbName default instead got %s", conn.dbName)
	}
}
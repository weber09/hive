package hive

import (
	"testing"
)

func TestHiveDrive_Open(t *testing.T) {
	drive := new(hiveDrive)
	conn, _ := drive.Open("jdbc:hive2://localhost:10000/default")

	if conn == nil {
		t.Errorf("Expected connection but got nil")
	}


}

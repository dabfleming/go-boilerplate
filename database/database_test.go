package database

import (
	"testing"
)

func TestConnection(t *testing.T) {
	var myVal int
	var err error

	err = db.QueryRow("SELECT 1 as col").Scan(&myVal)
	if err != nil {
		t.Fatalf("Error in row.Scan(): %v", err)
	}

	if myVal != 1 {
		t.Error("Scanned value %v, expected 1.", myVal)
	}
}

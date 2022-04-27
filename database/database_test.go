package database

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestInitializeForTest(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "TestInitializeForTest"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitializeForTest()
		})
	}
}

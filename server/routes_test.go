package server

import "testing"

func Test_initRoutes(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "prueba#1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initRoutes()
		})
	}
}

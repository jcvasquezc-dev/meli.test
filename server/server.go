package server

import (
	"net/http"
)

func Create(addr string) *http.Server {
	initRoutes()

	return &http.Server{
		Addr: addr,
	}
}

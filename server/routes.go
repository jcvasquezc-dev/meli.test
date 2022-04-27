package server

import (
	"net/http"
)

func initRoutes() {
	http.HandleFunc("/", Index)

	http.HandleFunc("/mutant", VerifyDnaSequence)

	http.HandleFunc("/mutant/", VerifyDnaSequence)

	http.HandleFunc("/stats", GetStats)

	http.HandleFunc("/stats/", GetStats)
}

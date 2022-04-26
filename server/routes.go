package server

import (
	"net/http"
)

func initRoutes() {
	http.HandleFunc("/", index)

	http.HandleFunc("/mutant", verifyDnaSequence)

	http.HandleFunc("/mutant/", verifyDnaSequence)

	http.HandleFunc("/stats", getStats)

	http.HandleFunc("/stats/", getStats)
}

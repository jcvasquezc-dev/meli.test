package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"meli.test/business/xmen"
	"meli.test/database"
	"meli.test/dtos"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404: Method not found")
}

func verifyDnaSequence(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed %s", r.Method)
		return
	}

	dna := &dtos.Dna{}
	err := json.NewDecoder(r.Body).Decode(dna)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if database.FindDnaBySequence(dna.Sequence) != nil {
		log.Println("the given dna sequence already exists and won't be analyzed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = xmen.IsMutant(dna)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "reason -> %s", err.Error())
		return
	}

	database.AddDna(dna)

	if !dna.IsMutant {
		w.WriteHeader(http.StatusForbidden)
	}
}

func getStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	stats := database.GetDnaStats()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

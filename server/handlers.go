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

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404: Method not found")
}

func VerifyDnaSequence(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("Method not allowed /mutant - ", r.Method)
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
		log.Println("reason -> ", err.Error())
		return
	}

	database.AddDna(dna)

	if !dna.IsMutant {
		w.WriteHeader(http.StatusForbidden)
	}
}

func GetStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("Method not allowed /stats - ", r.Method)
		return
	}

	stats := database.GetDnaStats()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

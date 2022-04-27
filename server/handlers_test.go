package server

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"meli.test/database"
)

type HttpRequestResult struct {
	Body       string
	Method     string
	StatusCode int
}

func TestIndex(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "prueba#1",
			args: args{
				r: httptest.NewRequest(http.MethodGet, "https://fakeurl.com", nil),
				w: httptest.NewRecorder(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Index(tt.args.w, tt.args.r)
		})
	}
}

var dt_1 = []HttpRequestResult{
	{ // empty structure
		Body:       `{ "dna": [] }`,
		Method:     http.MethodPost,
		StatusCode: http.StatusBadRequest,
	},
	{ // no structure present
		Body:       ``,
		Method:     http.MethodPost,
		StatusCode: http.StatusBadRequest,
	},
	{ // human dna
		Body:       `{ "dna": ["ATGCGA","CAGTGC","TTATGT","AGAACG"] }`,
		Method:     http.MethodPost,
		StatusCode: http.StatusForbidden,
	},
	{ // repeated human dna
		Body:       `{ "dna": ["ATGCGA","CAGTGC","TTATGT","AGAACG"] }`,
		Method:     http.MethodPost,
		StatusCode: http.StatusBadRequest,
	},
	{ // mutant dna
		Body:       `{ "dna": ["ATGCGA","CAGTGC","TTATGT","AGAACG","ATGCGA","CAGTGC","TTATGT","AGAACG"] }`,
		Method:     http.MethodPost,
		StatusCode: http.StatusOK,
	},
	{ // repeated mutant dna
		Body:       `{ "dna": ["ATGCGA","CAGTGC","TTATGT","AGAACG","ATGCGA","CAGTGC","TTATGT","AGAACG"] }`,
		Method:     http.MethodPost,
		StatusCode: http.StatusBadRequest,
	},
	{ // malformed dimensions by columns
		Body:       `{ "dna": ["ATGCGA","CAGGC","TTAT","AGAACGTTTT"] }`,
		Method:     http.MethodPost,
		StatusCode: http.StatusBadRequest,
	},
	{ // calling http method not allowed
		Body:       `{ "dna": ["ATGCGA","CAGTGC","TTATGT","AGAACG"] }`,
		Method:     http.MethodPut,
		StatusCode: http.StatusMethodNotAllowed,
	},
}

func TestVerifyDnaSequence(t *testing.T) {
	database.InitializeForTest()

	for _, httprequest := range dt_1 {
		req := httptest.NewRequest(httprequest.Method, "https://fakeurl.com", bytes.NewBuffer([]byte(httprequest.Body)))
		w := httptest.NewRecorder()

		VerifyDnaSequence(w, req)
		res := w.Result()
		//body, _ := ioutil.ReadAll(resp.Body)
		if res.StatusCode != httprequest.StatusCode {
			t.Fatal("test fail - http code expected ", httprequest.StatusCode, ", but got ", res.StatusCode)
		}
	}
}

var dt_2 = []HttpRequestResult{
	{ // allowed method
		Body:       "",
		Method:     http.MethodGet,
		StatusCode: http.StatusOK,
	},
	{ // not allowed method
		Body:       "",
		Method:     http.MethodOptions,
		StatusCode: http.StatusMethodNotAllowed,
	},
}

func TestGetStats(t *testing.T) {
	database.InitializeForTest()

	for _, httprequest := range dt_2 {
		req := httptest.NewRequest(httprequest.Method, "https://fakeurl.com", nil)
		w := httptest.NewRecorder()

		GetStats(w, req)
		res := w.Result()
		body, _ := ioutil.ReadAll(res.Body)

		log.Printf("%s", string(body))

		if res.StatusCode != httprequest.StatusCode {
			t.Fatal("test fail - http code expected ", httprequest.StatusCode, ", but got ", res.StatusCode)
		}
	}
}

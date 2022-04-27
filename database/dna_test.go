package database

import (
	"reflect"
	"testing"

	dto "meli.test/dtos"
)

func TestAddDna(t *testing.T) {
	type args struct {
		dna *dto.Dna
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "TestAddDna#1", args: args{dna: &dto.Dna{Sequence: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAACG"}, IsMutant: true}}},
		{name: "TestAddDna#2", args: args{dna: &dto.Dna{Sequence: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAACG"}, IsMutant: true}}},
	}
	InitializeForTest()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddDna(tt.args.dna)
		})
	}
}

func TestFindDnaBySequence(t *testing.T) {
	type args struct {
		sequence []string
	}
	tests := []struct {
		name string
		args args
		want *dto.Dna
	}{
		{name: "TestFindDnaBySequence#1", args: args{sequence: []string{}}, want: nil},
		{name: "TestFindDnaBySequence#2", args: args{sequence: []string{"AAAAA"}}, want: &dto.Dna{Id: 1, Sequence: []string{"AAAAA"}, IsMutant: true}},
	}
	InitializeForTest()
	AddDna(&dto.Dna{Sequence: []string{"AAAAA"}, IsMutant: true})
	AddDna(&dto.Dna{Sequence: []string{"ATATAT"}, IsMutant: false})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDnaBySequence(tt.args.sequence); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindDnaBySequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDnaStats(t *testing.T) {
	tests := []struct {
		name string
		want *dto.DnaStats
	}{
		{name: "TestGetDnaStats#1", want: &dto.DnaStats{}},
	}
	InitializeForTest()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDnaStats(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDnaStats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateRatio(t *testing.T) {
	type args struct {
		mutantsQty int
		humansQty  int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "Test_calculateRatio#1", args: args{mutantsQty: 0, humansQty: 0}, want: 0},
		{name: "Test_calculateRatio#1", args: args{mutantsQty: 0, humansQty: 50}, want: 0},
		{name: "Test_calculateRatio#1", args: args{mutantsQty: 1, humansQty: 10}, want: 0.1},
		{name: "Test_calculateRatio#1", args: args{mutantsQty: 20, humansQty: 10}, want: 2},
		{name: "Test_calculateRatio#1", args: args{mutantsQty: 50, humansQty: 0}, want: 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateRatio(tt.args.mutantsQty, tt.args.humansQty); got != tt.want {
				t.Errorf("calculateRatio() = %v, want %v", got, tt.want)
			}
		})
	}
}

package xmen

import (
	"reflect"
	"testing"

	dto "meli.test/dtos"
)

var defaultFormatedDna = [][]string{ // 6x6
	[]string{"A", "T", "G", "C", "G", "A"},
	[]string{"C", "A", "G", "T", "G", "C"},
	[]string{"T", "T", "A", "T", "G", "T"},
	[]string{"A", "G", "A", "A", "G", "G"},
	[]string{"C", "C", "C", "C", "T", "A"},
	[]string{"T", "C", "A", "C", "T", "G"},
}

func TestIsMutant(t *testing.T) {
	type args struct {
		dna *dto.Dna
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "TestIsMutant#1", args: args{&dto.Dna{Sequence: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}}}, want: true, wantErr: false},
		{name: "TestIsMutant#2", args: args{&dto.Dna{Sequence: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGXAGG", "CCCCTA", "TCACTG"}}}, want: false, wantErr: true},
		{name: "TestIsMutant#3", args: args{&dto.Dna{Sequence: []string{"ATGCGA", "CAGGC", "TTATGT", "AGXAGG", "CCCCTA", "TCACTG"}}}, want: false, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsMutant(tt.args.dna)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsMutant() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsMutant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFormatedDna(t *testing.T) {
	type args struct {
		dnaSequence []string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{name: "Test_getFormatedDna#1", args: args{dnaSequence: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}}, want: defaultFormatedDna, wantErr: false},
		{name: "Test_getFormatedDna#2", args: args{dnaSequence: []string{"ATGCGA", "CAGTGC"}}, want: [][]string{[]string{"A", "T", "G", "C", "G", "A"}, []string{"C", "A", "G", "T", "G", "C"}}, wantErr: false},
		{name: "Test_getFormatedDna#3", args: args{dnaSequence: []string{}}, want: nil, wantErr: true},
		{name: "Test_getFormatedDna#4", args: args{dnaSequence: []string{"ATGCGAA", "CAGTGC"}}, want: nil, wantErr: true},
		{name: "Test_getFormatedDna#5", args: args{dnaSequence: []string{"ATA", "CGC"}}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFormatedDna(tt.args.dnaSequence)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFormatedDna() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFormatedDna() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasValidElements(t *testing.T) {
	type args struct {
		dnaSequence []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test_hasValidElements#1", args: args{dnaSequence: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG"}}, want: true},
		{name: "Test_hasValidElements#2", args: args{dnaSequence: []string{"ATGCGA", "CAGTGC", "TTXTGT", "AGAAGG"}}, want: false},
		{name: "Test_hasValidElements#3", args: args{dnaSequence: []string{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasValidElements(tt.args.dnaSequence); got != tt.want {
				t.Errorf("hasValidElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_analyzeSequenceAsString(t *testing.T) {
	type args struct {
		subDnaSequence string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test_analyzeSequenceAsString#1", args: args{subDnaSequence: "AAATTTGGGCCC"}, want: 0},
		{name: "Test_analyzeSequenceAsString#2", args: args{subDnaSequence: "AAAAATGC"}, want: 1},
		{name: "Test_analyzeSequenceAsString#3", args: args{subDnaSequence: "AAAAAAAAA"}, want: 2},
		{name: "Test_analyzeSequenceAsString#4", args: args{subDnaSequence: "AAAATTTTCCCC"}, want: 2},
		{name: "Test_analyzeSequenceAsString#5", args: args{subDnaSequence: "AAAATTTTCCCCGGGG"}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := analyzeSequenceAsString(tt.args.subDnaSequence); got != tt.want {
				t.Errorf("analyzeSequenceAsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasMinimumMatchesRequired(t *testing.T) {
	type args struct {
		matchesFound int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test_hasMinimumMatchesRequired#1", args: args{matchesFound: 0}, want: false},
		{name: "Test_hasMinimumMatchesRequired#2", args: args{matchesFound: 1}, want: false},
		{name: "Test_hasMinimumMatchesRequired#3", args: args{matchesFound: 2}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasMinimumMatchesRequired(tt.args.matchesFound); got != tt.want {
				t.Errorf("hasMinimumMatchesRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_doDnaAnalysis(t *testing.T) {
	type args struct {
		formatedDna [][]string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test_doDnaAnalysis#1",
			args: args{ // 6x6 with 2 matches
				formatedDna: defaultFormatedDna,
			},
			want: 2,
		},
		{
			name: "Test_doDnaAnalysis#2",
			args: args{
				formatedDna: [][]string{ // 6x6 with not matches
					[]string{"A", "A", "T", "T", "C", "A"},
					[]string{"A", "G", "C", "C", "G", "C"},
					[]string{"T", "A", "T", "T", "A", "T"},
					[]string{"A", "A", "G", "G", "A", "A"},
					[]string{"C", "A", "C", "C", "A", "C"},
					[]string{"G", "C", "T", "T", "C", "G"},
				},
			},
			want: 0,
		},
		{
			name: "Test_doDnaAnalysis#3",
			args: args{
				formatedDna: [][]string{ // 6x3 with not matches because the size
					[]string{"A", "A", "T"},
					[]string{"A", "G", "C"},
					[]string{"T", "A", "T"},
					[]string{"A", "A", "T"},
					[]string{"A", "G", "C"},
					[]string{"T", "A", "T"},
				},
			},
			want: 0,
		},
		{
			name: "Test_doDnaAnalysis#4",
			args: args{
				formatedDna: [][]string{ // 6x3 with not matches because the size
					[]string{"A", "A", "T"},
					[]string{"A", "G", "C"},
					[]string{"T", "A", "T"},
					[]string{"A", "A", "T"},
					[]string{"A", "A", "C"},
					[]string{"T", "A", "T"},
				},
			},
			want: 1,
		},
		{
			name: "Test_doDnaAnalysis#5",
			args: args{
				formatedDna: [][]string{ // 3x6 with not matches because the size
					[]string{"A", "A", "T", "T", "C", "A"},
					[]string{"A", "G", "C", "C", "G", "C"},
					[]string{"T", "A", "T", "T", "A", "T"},
				},
			},
			want: 0,
		},
		{
			name: "Test_doDnaAnalysis#6",
			args: args{
				formatedDna: [][]string{ // 3x6 with not matches because the size
					[]string{"A", "A", "T", "T", "C", "A"},
					[]string{"A", "G", "C", "C", "G", "C"},
					[]string{"T", "A", "T", "T", "T", "T"},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doDnaAnalysis(tt.args.formatedDna); got != tt.want {
				t.Errorf("doDnaAnalysis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_horizontalSearch(t *testing.T) {
	type args struct {
		formatedDna  [][]string
		matchesFound int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test_horizontalSearch#1",
			args: args{ // 6x6 with 1 match
				formatedDna:  defaultFormatedDna,
				matchesFound: 0,
			},
			want: 1,
		},
		{
			name: "Test_horizontalSearch#2",
			args: args{ // 6x6 with 1 match
				formatedDna:  defaultFormatedDna,
				matchesFound: 1,
			},
			want: 2,
		},
		{
			name: "Test_horizontalSearch#3",
			args: args{
				formatedDna: [][]string{ // 6x6 with not matches
					[]string{"A", "A", "T", "T", "C", "A"},
					[]string{"A", "G", "C", "C", "G", "C"},
					[]string{"T", "A", "T", "T", "A", "T"},
					[]string{"A", "A", "G", "G", "A", "A"},
					[]string{"C", "A", "C", "C", "A", "C"},
					[]string{"G", "C", "T", "T", "C", "G"},
				},
				matchesFound: 1,
			},
			want: 1,
		},
		{
			name: "Test_horizontalSearch#4",
			args: args{
				formatedDna: [][]string{ // 6x3 with not matches because the size
					[]string{"A", "A", "T"},
					[]string{"A", "G", "C"},
					[]string{"T", "A", "T"},
					[]string{"A", "A", "T"},
					[]string{"A", "G", "C"},
					[]string{"T", "A", "T"},
				},
				matchesFound: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := horizontalSearch(tt.args.formatedDna, tt.args.matchesFound); got != tt.want {
				t.Errorf("horizontalSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_verticalSearch(t *testing.T) {
	type args struct {
		formatedDna  [][]string
		matchesFound int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test_verticalSearch#1",
			args: args{ // 6x6 with 1 match
				formatedDna:  defaultFormatedDna,
				matchesFound: 0,
			},
			want: 1,
		},
		{
			name: "Test_verticalSearch#2",
			args: args{ // 6x6 with 1 match
				formatedDna:  defaultFormatedDna,
				matchesFound: 1,
			},
			want: 2,
		},
		{
			name: "Test_verticalSearch#3",
			args: args{
				formatedDna: [][]string{ // 6x6 with not matches
					[]string{"A", "A", "T", "T", "C", "A"},
					[]string{"A", "G", "C", "C", "G", "C"},
					[]string{"T", "A", "T", "T", "A", "T"},
					[]string{"A", "A", "G", "G", "A", "A"},
					[]string{"C", "A", "C", "C", "A", "C"},
					[]string{"G", "C", "T", "T", "C", "G"},
				},
				matchesFound: 1,
			},
			want: 1,
		},
		{
			name: "Test_verticalSearch#4",
			args: args{
				formatedDna: [][]string{ // 3x6 with not matches because the size
					[]string{"A", "A", "T", "T", "C", "A"},
					[]string{"A", "G", "C", "C", "G", "C"},
					[]string{"T", "A", "T", "T", "A", "T"},
				},
				matchesFound: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verticalSearch(tt.args.formatedDna, tt.args.matchesFound); got != tt.want {
				t.Errorf("verticalSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_obliqueSearch(t *testing.T) {
	type args struct {
		formatedDna  [][]string
		matchesFound int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test_obliqueSearch#1",
			args: args{ // 6x6 with 1 match
				formatedDna:  defaultFormatedDna,
				matchesFound: 0,
			},
			want: 1,
		},
		{
			name: "Test_obliqueSearch#2",
			args: args{ // 6x6 with 1 match
				formatedDna:  defaultFormatedDna,
				matchesFound: 1,
			},
			want: 2,
		},
		{
			name: "Test_obliqueSearch#3",
			args: args{
				formatedDna: [][]string{ // 6x6 with not matches
					[]string{"A", "A", "T", "T", "C", "A"},
					[]string{"A", "G", "C", "C", "G", "C"},
					[]string{"T", "A", "T", "T", "A", "T"},
					[]string{"A", "A", "G", "G", "A", "A"},
					[]string{"C", "A", "C", "C", "A", "C"},
					[]string{"G", "C", "T", "T", "C", "G"},
				},
				matchesFound: 1,
			},
			want: 1,
		},
		{
			name: "Test_obliqueSearch#4",
			args: args{
				formatedDna: [][]string{ // 3x6 with not matches because the size
					[]string{"A", "A", "T", "T", "C", "A"},
					[]string{"A", "G", "C", "C", "G", "C"},
					[]string{"T", "A", "T", "T", "A", "T"},
				},
				matchesFound: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := obliqueSearch(tt.args.formatedDna, tt.args.matchesFound); got != tt.want {
				t.Errorf("obliqueSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

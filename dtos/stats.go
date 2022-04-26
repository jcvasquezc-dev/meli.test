package dtos

type DnaStats struct {
	MutantsQty int     `json:"count_mutant_dna"`
	HumansQty  int     `json:"count_human_dna"`
	Ratio      float64 `json:"ratio"`
}

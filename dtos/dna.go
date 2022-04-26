package dtos

type Dna struct {
	Id       int
	Sequence []string `json:"dna,omitempty"`
	IsMutant bool
}

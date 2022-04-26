package xmen

import (
	"errors"
	"fmt"
	"strings"

	dto "meli.test/dtos"
)

const allowedDnaValues = "ATCG"  // set of allowed values to be analyzed
const matchSequenceCount = 4     // total of occurrences that should be reached to be considered a valid match
const showMatrixLog = false      // print a multidimensaional array representation of the dna sequence in the console
const minimumMatchesRequired = 2 // quantity of minimum occurrences that should be reached for the dna sequence to be considered as mutant

// IsMutantTestMethod is a public method that determine if a dna sequence comes from a human or mutant.
// This method should be used for testing purpose.
// It returns true if the dna sequence belongs to a mutant or a human; return false otherwise. Also return an error if exists.
func IsMutant(dna *dto.Dna) (bool, error) {
	if !hasValidElements(dna.Sequence) {
		return false, errors.New("the DNA sequence comes from an superior species because has unrecognized values")
	}

	formatedDna, er := getFormatedDna(dna.Sequence)

	// dna is corrupted by size (array dimensions), or for absence of some elements
	if er != nil {
		return false, er
	}

	matchesFound := doDnaAnalysis(formatedDna)
	dna.IsMutant = hasMinimumMatchesRequired(matchesFound)

	return dna.IsMutant, nil
}

// IsMutantTestMethod is a public method that determine if a dna sequence comes from a human or mutant.
// This method should be used for testing purpose.
// It returns true if the dna sequence belongs to a mutant or a human; return false otherwise. Also return the number of matches founds. Also return an error if exists.
func IsMutantTestMethod(dna *dto.Dna) (bool, int, error) {
	if !hasValidElements(dna.Sequence) {
		return false, 0, errors.New("the DNA sequence comes from an superior species because has unrecognized values")
	}

	formatedDna, er := getFormatedDna(dna.Sequence)

	// dna is corrupted by size (array dimensions), or for absence of some elements
	if er != nil {
		return false, 0, er
	}

	matchesFound := doDnaAnalysis(formatedDna)
	dna.IsMutant = hasMinimumMatchesRequired(matchesFound)

	return dna.IsMutant, matchesFound, nil
}

// getFormatedDna is a private method that converts the dna sequence received into a formated dimensional matrix.
// All the validations needed for the dna sequence are mades here. When some validations are mades and fails, a error is returned.
// It returns a multidimensional array representation of the original dna sequence.
func getFormatedDna(dnaSequence []string) ([][]string, error) {
	var formatedDna [][]string

	// validating that DNA sequence has at least one element
	if len(dnaSequence) == 0 {
		return nil, errors.New("no DNA found")
	}

	if len(dnaSequence) > 1 {
		nitrogenBaseSizeReference := len(dnaSequence[0])

		for i := 0; i < len(dnaSequence); i++ {
			if len(dnaSequence[i]) != nitrogenBaseSizeReference {
				return nil, errors.New("the DNA provided is corrupt")
			}

			nitrogenBase := []string{} // row

			if showMatrixLog {
				fmt.Printf("[")
			}

			for _, val := range dnaSequence[i] {
				nitrogenBase = append(nitrogenBase, strings.ToUpper(string(val)))

				if showMatrixLog {
					fmt.Printf(" %s ", strings.ToUpper(string(val)))
				}
			}

			if showMatrixLog {
				fmt.Println("]")
			}

			formatedDna = append(formatedDna, nitrogenBase)
		}

		// if both of the dimensions fo the dnaSequence (row/column) are less than matchSequenceCount
		if len(dnaSequence) < matchSequenceCount && nitrogenBaseSizeReference < matchSequenceCount {
			return nil, errors.New("the DNA found is too premature to be analyzed")
		}
	}

	return formatedDna, nil
}

// hasValidElements is a private method that evaluate all the dna sequence in of invalid elements.
// It returns true if the complete dna sequence only contains valid elements; return false otherwise.
func hasValidElements(dnaSequence []string) bool {
	dnaSequenceValues := strings.ToUpper(strings.Join(dnaSequence, ""))

	for _, val := range allowedDnaValues {
		dnaSequenceValues = strings.ReplaceAll(dnaSequenceValues, string(val), "")
	}

	return len(dnaSequenceValues) == 0
}

// analyzeSequenceAsString is a private method that compare a sub-dnaSequence with the allowedDnaValues.
// It returns the number of matches between sub-dnaSequence and repeated elements of allowedDnaValues.
func analyzeSequenceAsString(subDnaSequence string) int {
	matchesFound := 0

	for _, val := range allowedDnaValues {
		validMatch := strings.Repeat(string(val), matchSequenceCount)
		matchesFound += strings.Count(subDnaSequence, validMatch)

		if hasMinimumMatchesRequired(matchesFound) {
			break
		}
	}

	return matchesFound
}

// hasMinimumMatchesRequired is a private method that compare a number of matches with the minimumMatchesRequired.
// It returns true if the matchesFound is equal or higher than minimumMatchesRequired; returns false otherwise.
func hasMinimumMatchesRequired(matchesFound int) bool {
	return matchesFound >= minimumMatchesRequired
}

// doDnaAnalysis is a private method that call the implementations of: obliqueSearch, horizontalSearch and verticalSearch.
// It returns the number of minimum matches required to be considerer as mutant (minimumMatchesRequired).
func doDnaAnalysis(formatedDna [][]string) int {
	matchesFound := 0

	matchesFound = horizontalSearch(formatedDna, matchesFound)

	if !hasMinimumMatchesRequired(matchesFound) {
		matchesFound = verticalSearch(formatedDna, matchesFound)

		if !hasMinimumMatchesRequired(matchesFound) {
			matchesFound = obliqueSearch(formatedDna, matchesFound)
		}
	}

	return matchesFound
}

// horizontalSearch is a private method that search horizontally (-) patrons that match with allowedDnaValues values.
// It returns the number of matches with a sequence of allowedDnaValues characters (matchSequenceCount in a row).
func horizontalSearch(formatedDna [][]string, matchesFound int) int {
	columnCount := len(formatedDna[0])

	if columnCount < matchSequenceCount {
		return 0
	}

	for _, row := range formatedDna {
		dnaSequence := strings.Join(row, "")
		matchesFound += analyzeSequenceAsString(dnaSequence)

		if hasMinimumMatchesRequired(matchesFound) {
			break
		}
	}

	return matchesFound
}

// verticalSearch is a private method that search vertically (|) patrons that match with allowedDnaValues values.
// It returns the number of matches with a sequence of allowedDnaValues characters (matchSequenceCount in a row).
func verticalSearch(formatedDna [][]string, matchesFound int) int {
	rowCount := len(formatedDna)
	columnCount := len(formatedDna[0])

	if rowCount < matchSequenceCount {
		return 0
	}

	for j := 0; j < columnCount; j++ {
		dnaSequence := ""

		for i := 0; i < rowCount; i++ {
			dnaSequence += formatedDna[i][j]
		}

		matchesFound += analyzeSequenceAsString(dnaSequence)

		if hasMinimumMatchesRequired(matchesFound) {
			break
		}
	}

	return matchesFound
}

// obliqueSearch is a private method that search obliquely patrons that match with allowedDnaValues values.
// The oblique search is done searching from left to rigth (\), and searching from right to left (/) (both from top to bottom)
// It returns the number of matches with a sequence of allowedDnaValues characters (matchSequenceCount in a row).
func obliqueSearch(formatedDna [][]string, matchesFound int) int {
	rowCount := len(formatedDna)
	columnCount := len(formatedDna[0])

	if rowCount < matchSequenceCount || columnCount < matchSequenceCount {
		return 0
	}

	colummnTopIndex := columnCount - matchSequenceCount
	var (
		i int = rowCount - matchSequenceCount // rowStartIndex
		j int = 0
	)

	for j < colummnTopIndex {
		dnaSequenceL2R := "" // Oblique left to rigth from top
		dnaSequenceR2L := "" // Oblique rigth to left from top
		r := i
		c := j

		for r < rowCount && c < columnCount {
			dnaSequenceL2R = dnaSequenceL2R + formatedDna[r][c]
			dnaSequenceR2L = dnaSequenceR2L + formatedDna[r][columnCount-1-c]
			r++
			c++
		}

		matchesFound += analyzeSequenceAsString(dnaSequenceL2R)

		if hasMinimumMatchesRequired(matchesFound) {
			break
		}

		matchesFound += analyzeSequenceAsString(dnaSequenceR2L)

		if hasMinimumMatchesRequired(matchesFound) {
			break
		}

		if i == 0 {
			j++
		} else {
			i--
		}
	}

	return matchesFound
}

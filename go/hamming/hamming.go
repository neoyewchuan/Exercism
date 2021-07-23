package hamming

import "errors"

// Calculate the hamming distance between to strands of DNA
func Distance(a, b string) (int, error) {
	// Both DNA strains must be of the same length, otherwise raise exception.
	if len(a) != len(b) {
		return -1, errors.New("Invalid DNA strands pair.")
	}
	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}

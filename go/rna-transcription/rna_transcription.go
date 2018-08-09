// Package strand provides a library for manipulating DNA an RNA strands.
package strand

// ToRNA converts a DNA strand into an RNA strand, if successful,
// the resulting RNA strand is returned.
func ToRNA(dna string) string {
	rna := make([]byte, len(dna))
	for i, r := range dna {
		switch r {
		case 'G':
			rna[i] = 'C'
		case 'C':
			rna[i] = 'G'
		case 'T':
			rna[i] = 'A'
		case 'A':
			rna[i] = 'U'
		}
	}
	return string(rna)
}

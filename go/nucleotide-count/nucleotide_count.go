// Package dna provides a library for working with dna.
package dna

import "errors"

var ErrInvalidNucleotides = errors.New("Invalid Nucleotides in DNA.")

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// NewHistogram is the constructor of the Histogram type.
func NewHistogram() Histogram {
	return Histogram(map[rune]int{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	})
}

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := NewHistogram()
	for _, r := range d {
		if r != 'A' && r != 'C' && r != 'G' && r != 'T' {
			return h, ErrInvalidNucleotides
		}
		h[r]++
	}
	return h, nil
}

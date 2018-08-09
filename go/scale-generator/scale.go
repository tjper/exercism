// Package scale provides a library for interacting with scales in Western music.
package scale

import "strings"

const ChromaticLength = 12

// Scale generates a musical scale based on the tonic and interval, if successful,
// the musical scale is returned.
func Scale(tonic, interval string) []string {
	var scaleSet []string
	switch tonic {
	case "C", "G", "D", "A", "E", "B", "F#", "a", "e", "b", "f#", "c#", "g#", "d#":
		scaleSet = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		scaleSet = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
	}

	start := tonicIndex(tonic, scaleSet)
	scale := make([]string, 0)

	if interval == "" {
		for i := start; i < (start + ChromaticLength); i++ {
			scale = append(scale, scaleSet[(i%len(scaleSet))])
		}
	} else {
		index := start
		for _, note := range interval {
			scale = append(scale, scaleSet[index])
			switch note {
			case 'm':
				index = (index + 1) % len(scaleSet)
			case 'M':
				index = (index + 2) % len(scaleSet)
			case 'A':
				index = (index + 3) % len(scaleSet)
			}
		}
	}

	return scale

}

// tonicIndex determines the index of tonic within a Chromatic Scale.
// Returns index
func tonicIndex(tonic string, chromScale []string) int {
	tonic = strings.Title(tonic)
	var index int
	for i, note := range chromScale {
		if tonic == note {
			index = i
			break
		}
	}
	return index
}

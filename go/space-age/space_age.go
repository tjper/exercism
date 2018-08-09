// Package space provides transforms of age for different planets.
package space

// Constants are number of seconds in a year for the respective planet.
const (
	EarthSecPerYear   = 31557600.00
	MercurySecPerYear = EarthSecPerYear * .2408467
	VenusSecPerYear   = EarthSecPerYear * .61519726
	MarsSecPerYear    = EarthSecPerYear * 1.8808158
	JupiterSecPerYear = EarthSecPerYear * 11.862615
	SaturnSecPerYear  = EarthSecPerYear * 29.447498
	UranusSecPerYear  = EarthSecPerYear * 84.016846
	NeptuneSecPerYear = EarthSecPerYear * 164.79132
)

// Planet is a data type representing each planet in the Solar System.
type Planet string

const (
	Earth   Planet = "Earth"
	Mercury Planet = "Mercury"
	Venus   Planet = "Venus"
	Mars    Planet = "Mars"
	Jupiter Planet = "Jupiter"
	Saturn  Planet = "Saturn"
	Uranus  Planet = "Uranus"
	Neptune Planet = "Neptune"
)

// Age determines the age in planetary years of an object, if successful,
// a the an age in years is returned.
func Age(s float64, planet Planet) float64 {
	switch planet {
	case Earth:
		return s / EarthSecPerYear
	case Mercury:
		return s / MercurySecPerYear
	case Venus:
		return s / VenusSecPerYear
	case Mars:
		return s / MarsSecPerYear
	case Jupiter:
		return s / JupiterSecPerYear
	case Saturn:
		return s / SaturnSecPerYear
	case Uranus:
		return s / UranusSecPerYear
	}
	return s / NeptuneSecPerYear
}

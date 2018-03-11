// package conv provides converters for temperature, distance and weight ex 2.2
// write a general purpose conversion program analogous to cf that reads numbers
// from it's command line arguments or from the standard input if there are no
// arguments, and convert each number into units like temperature in celcius and
// fahrenheit, length in feet and meters, weight in pounds and kilograms, and
// the like.

package conv

// stringer creates Mode -> String function.
//go:generate stringer -type=Mode
type Mode int

const (
	_ Mode = iota
	Temperature
	Distance
	Weight
)

// Modes is a map of string to Mode iota.
var Modes = map[string]Mode {
	"temperature": Temperature,
	"distance": Distance,
	"weight": Weight,
}

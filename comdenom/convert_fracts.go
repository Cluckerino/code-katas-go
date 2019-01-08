package comdenom

import (
	"fmt"
	"strings"
)

// ConvertFracts turns a slice of fractions and prints them in LCM form.
func ConvertFracts(fracts [][]int) string {
	denoms := []int{}
	for _, fract := range fracts {
		// Make sure the inputs are from reduced fractions!!!
		gcf := Gcf(fract[0], fract[1])
		denom := fract[1] / gcf
		denoms = append(denoms, denom)
	}

	lcd := MultiLcm(denoms)
	stringFracts := []string{}
	for _, fract := range fracts {
		// Calculate the numerator for the fraction with the LCD.
		convNum := (fract[0] * lcd) / fract[1]
		stringFract := fmt.Sprintf("(%v,%v)", convNum, lcd)
		stringFracts = append(stringFracts, stringFract)
	}

	return strings.Join(stringFracts, "")
}

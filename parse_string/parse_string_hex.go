package parse_string

import "math"

// ParseStringDecimal erwartet einen String, der eine Hexadezimalzahl repräsentiert, und liefert die zugehörige Zahl.
// Ist der String kein gültiger Wert, wird -1 zurückgegeben.
func ParseStringHexadecimal(s string) int {

	decimalNumber := 0
	PreviousBase := 16

	//Diese Stellen sind manchmal größer als 9, sodass sie einfach addiert werden müssen, um die Zahl zur Basis 10 darzustellen
	for j, el2 := range s {

		//decimalNumber wird mit iterierendem addieren vom Element an der Stelle j mit 16^ len(s) -1 -j errechnet, da das erste element des Strings das most significant digit ist
		decimalNumber += ParseDigit(string(el2)) * int(math.Pow(float64(PreviousBase), float64(len(s)-1-j)))
	}

	return decimalNumber
}

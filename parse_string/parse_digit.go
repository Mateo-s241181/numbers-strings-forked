package parse_string

// ParseDigit erwartet einen String von 0 bis 9 oder A bis F und liefert die zugehörige Zahl.
// Dabei gilt A=10, B=11, ..., F=15.
// Ist der String kein gültiger Wert, wird -1 zurückgegeben.
func ParseDigit(digit string) int {

	//Diese Strings sind zur Eingabe erlaubt

	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	expected_digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	//Wenn das Eingabedigit im String "digits" enthalten ist, soll der Wert an der gleichen Position aus der Liste "expected_digits" zurückgegeben werden

	for i, el := range digits {
		if digit == el {
			return expected_digits[i]
		}
	}

	//Falls das obige return Statement nicht erreicht wird, ist der Eingegebene String nicht in "digits" enthalten

	return -1
}

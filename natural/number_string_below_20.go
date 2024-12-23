package natural

// NumberStringBelow20 erwartet eine Zahl < 20 und liefert den zugehörigen String.
func NumberStringBelow20(n int) string {
	DigitStringBelowTwenty := ""
	thirdDigit := (n - (n % 100)) / 100

	switch n {

	case 0:
		DigitStringBelowTwenty = "null"
	case 1:
		DigitStringBelowTwenty = "eins"
	case 2:
		DigitStringBelowTwenty = "zwei"
	case 3:
		DigitStringBelowTwenty = "drei"
	case 4:
		DigitStringBelowTwenty = "vier"
	case 5:
		DigitStringBelowTwenty = "fünf"
	case 6:
		DigitStringBelowTwenty = "sechs"
	case 7:
		DigitStringBelowTwenty = "sieben"
	case 8:
		DigitStringBelowTwenty = "acht"
	case 9:
		DigitStringBelowTwenty = "neun"
	case 10:
		DigitStringBelowTwenty = "zehn"
	case 11:
		DigitStringBelowTwenty = "elf"
	case 12:
		DigitStringBelowTwenty = "zwölf"
	case 13:
		DigitStringBelowTwenty = "dreizehn"
	case 14:
		DigitStringBelowTwenty = "vierzehn"
	case 15:
		DigitStringBelowTwenty = "fünfzehn"
	case 16:
		DigitStringBelowTwenty = "sechzehn"
	case 17:
		DigitStringBelowTwenty = "siebzehn"
	case 18:
		DigitStringBelowTwenty = "achtzehn"
	case 19:
		DigitStringBelowTwenty = "neunzehn"

	}

	return DigitString100(thirdDigit) + DigitStringBelowTwenty
}

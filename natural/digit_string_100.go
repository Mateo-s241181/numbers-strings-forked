package natural

// DigitString100 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Hunderter-Stelle einer Zahl >= 21 vorkommen würde.
//
// Beispiel: Für 3 soll der String "dreihundert" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
func DigitString100(digit int) string {

	DigitAsString100 := ""
	Suffix100 := "hundert"

	switch digit {

	case 0:
		return ""
	case 1:
		DigitAsString100 = "ein"
	case 2:
		DigitAsString100 = "zwei"
	case 3:
		DigitAsString100 = "drei"
	case 4:
		DigitAsString100 = "vier"
	case 5:
		DigitAsString100 = "fünf"
	case 6:
		DigitAsString100 = "sechs"
	case 7:
		DigitAsString100 = "sieben"
	case 8:
		DigitAsString100 = "acht"
	case 9:
		DigitAsString100 = "neun"
	}
	return DigitAsString100 + Suffix100
}

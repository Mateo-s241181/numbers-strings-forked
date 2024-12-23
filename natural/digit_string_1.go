package natural

// DigitString1 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Einer-Stelle einer Zahl >= 21 vorkommen würde.
// Außerdem wird bei Ziffern != 0 das Wort "und" angehängt.
//
// Beispiel: Für 3 soll der String "dreiund" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
// Diese Funktion muss nur für den Normalfall (Zahlen >= 21) funktionieren.
func DigitString1(digit int) string {

	DigitAsString := ""
	placeholderString := "und"

	switch digit {
	case 1:
		DigitAsString = "ein"
	case 2:
		DigitAsString = "zwei"
	case 3:
		DigitAsString = "drei"
	case 4:
		DigitAsString = "vier"
	case 5:
		DigitAsString = "fünf"
	case 6:
		DigitAsString = "sechs"
	case 7:
		DigitAsString = "sieben"
	case 8:
		DigitAsString = "acht"
	case 9:
		DigitAsString = "neun"
	}

	if digit == 0 {
		return ""
	}

	return DigitAsString + placeholderString
}

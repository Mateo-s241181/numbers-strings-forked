package natural

// DigitString10 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Zehner-Stelle einer Zahl >= 21 vorkommen würde.
//
// Beispiel: Für 3 soll der String "dreißig" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
func DigitString10(digit int) string {

	DigitAsString10 := ""

	switch digit {

	case 0:
		return ""
	case 1:
		return ""
	case 2:
		DigitAsString10 = "zwanzig"
	case 3:
		DigitAsString10 = "dreißig"
	case 4:
		DigitAsString10 = "vierzig"
	case 5:
		DigitAsString10 = "fünfzig"
	case 6:
		DigitAsString10 = "sechzig"
	case 7:
		DigitAsString10 = "siebzig"
	case 8:
		DigitAsString10 = "achtzig"
	case 9:
		DigitAsString10 = "neunzig"
	}

	return DigitAsString10
}

func NumberBetweenTenAndTwenty(digitOne int) string {

	DigitAsStringTentoTwenty := ""

	switch digitOne {

	case 0:
		DigitAsStringTentoTwenty = "zehn"
	case 1:
		DigitAsStringTentoTwenty = "elf"
	case 2:
		DigitAsStringTentoTwenty = "zwölf"
	case 3:
		DigitAsStringTentoTwenty = "dreizehn"
	case 4:
		DigitAsStringTentoTwenty = "vierzehn"
	case 5:
		DigitAsStringTentoTwenty = "fünfzehn"
	case 6:
		DigitAsStringTentoTwenty = "sechzehn"
	case 7:
		DigitAsStringTentoTwenty = "siebzehn"
	case 8:
		DigitAsStringTentoTwenty = "achtzehn"
	case 9:
		DigitAsStringTentoTwenty = "neunzehn"

	}

	return DigitAsStringTentoTwenty
}

func NumberUnderTen(digit int) string {

	DigitAsStringUnderTen := ""

	switch digit {

	case 0:
		DigitAsStringUnderTen = "null"
	case 1:
		DigitAsStringUnderTen = "eins"
	case 2:
		DigitAsStringUnderTen = "zwei"
	case 3:
		DigitAsStringUnderTen = "drei"
	case 4:
		DigitAsStringUnderTen = "vier"
	case 5:
		DigitAsStringUnderTen = "fünf"
	case 6:
		DigitAsStringUnderTen = "sechs"
	case 7:
		DigitAsStringUnderTen = "sieben"
	case 8:
		DigitAsStringUnderTen = "acht"
	case 9:
		DigitAsStringUnderTen = "neun"
	}

	return DigitAsStringUnderTen
}

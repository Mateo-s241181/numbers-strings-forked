package natural

// NumberString6Digits erwartet eine Zahl 0 <= n <= 999999 und liefert den zugehörigen String.
func NumberString6Digits(n int) string {

	if n < 1000 {
		return NumberString3Digits(n)
	}

	if n >= 1000 {
		// Betrachtet nur die stellen vom Wert über 100
		m := (n - (n % 1000)) / 1000

		//Wenn alle Ziffern mit Wertigkeit < 1000, 0 sind, sollen Sie außen vor gelassen werden
		if m*1000 == n {
			return NumberString3Digits(m) + "tausend"
		}

		//Nutzt 3 Digit und Suffix -tausend
		return NumberString3Digits(m) + "tausend" + NumberString3Digits(n-(1000*m))
	}
	return ""
}

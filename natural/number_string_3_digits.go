package natural

// NumberString3Digits erwartet eine Zahl 0 <= n <= 999 und liefert den zugehörigen String.
func NumberString3Digits(n int) string {

	firstDigit := n % 10 //Least Significant Digit
	secondDigit := ((n - (n % 10)) % 100) / 10
	thirdDigit := (n - (n % 100)) / 100 //Most Significant Digit

	if secondDigit < 2 {
		//Ist die Zahl unter 100 kann sie Direkt in die Funktion eingesetzt werden (100er Stelle wird ansonsten als "null" ausgegeben)
		if n <= 100 {
			return NumberStringBelow20(n)
		}
		//Wenn die Zahl über 100 ist müssen die einzelnen Digits in die Jeweiligen Funktionen eingesetzt werden
		return DigitString100(thirdDigit) + NumberStringBelow20(10*secondDigit+firstDigit)
	}

	return NumberStringGreater20(n)
}

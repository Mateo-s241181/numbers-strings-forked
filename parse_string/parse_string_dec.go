package parse_string

import "strconv"

// ParseStringDecimal erwartet einen String, der eine Dezimalzahl repräsentiert, und liefert die zugehörige Zahl.
// Ist der String kein gültiger Wert, wird -1 zurückgegeben.
func ParseStringDecimal(s string) int {

	//Äquvalent zu ParseInt(s, 10, 32) 10 ist hierbei die Basis und 32 die Bitsize
	//s wird zu einem Int konvertiert und der Wert wird zurückgegeben, wenn der Error == nil ist
	if s, err := strconv.Atoi(s); err == nil {
		return s
	}

	//Wenn err != nil gilt, soll -1 als Fehlermeldung zurückgegeben werden
	return -1
}

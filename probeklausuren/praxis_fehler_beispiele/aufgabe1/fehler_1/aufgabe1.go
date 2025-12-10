package aufgabe1

// LongestAbc erwartet eine Liste von Strings und liefert
// das längste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
func LongestAbc(list []string) string {

	if len(list) == 0 {
		return ""
	}
	longestLen := ""

	for _, val := range list {

		if len(val) >= 3 && val[:3] == "abc" {
			if len(val) > len(longestLen) {
				longestLen = val
			}

		}

	}
	return longestLen
}

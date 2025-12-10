package aufgabe25

/* AUFGABENSTELLUNG:
 *
 * Schreiben Sie eine Funktion FilterShortStrings, die aus einer Liste von Strings
 * alle Strings entfernt, deren Länge < minLen ist.
 *
 * Beispiele:
 *   FilterShortStrings([]string{"hi","hello","a","abc"}, 3) -> ["hello","abc"]
 *   FilterShortStrings([]string{}, 5)                       -> []
 */

func FilterShortStrings(list []string, minLen int) []string {
	result := []string{}
	if len(list) == 0 {
		return result
	}

	for _, val := range list {
		if len(val) < minLen {

		} else {
			result = append(result, val)
		}
	}

	return result
}

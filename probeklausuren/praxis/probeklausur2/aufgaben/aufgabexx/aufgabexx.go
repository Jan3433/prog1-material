package aufgabeXX

/* AUFGABENSTELLUNG:
 *
 * RemoveChar entfernt alle Vorkommen des Zeichens c aus dem String s.
 *
 * Beispiele:
 *   RemoveChar("banana", 'a') -> "bnn"
 *   RemoveChar("hello", 'l')  -> "heo"
 *   RemoveChar("aaaa", 'a')   -> ""
 */

func RemoveChar(s string, c byte) string {
	result := ""
	if s == "" {
		return result
	}
	for i := 0; i < len(s); i++ {
		if s[i] != c {
			result += string(s[i])
		}
	}

	return result
}

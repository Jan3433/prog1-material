package aufgabe26

/* AUFGABENSTELLUNG:
 *
 * ReplaceAll ersetzt in einem String s alle Vorkommen eines Zeichens old
 * durch das Zeichen new.
 *
 * Beispiele:
 *   ReplaceAll("banana", 'a', 'o') -> "bonono"
 *   ReplaceAll("test", 't', 'x')  -> "xesx"
 */

func ReplaceAll(s string, old byte, new byte) string {
	result := ""
	if s == "" {
		return result
	}

	for i := 0; i < len(s); i++ {
		if s[i] == old {
			result += string(new)
		}
		if s[i] != old {
			result += string(s[i])
		}

	}
	return result
}

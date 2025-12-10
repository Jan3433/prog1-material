package aufgabe27

/* AUFGABENSTELLUNG:
 *
 * Schreiben Sie eine Funktion FirstIndexOf, die den Index des ersten
 * Auftretens eines Strings needle in einer Liste list liefert.
 *
 * Wird needle nicht gefunden, geben Sie -1 zurück.
 *
 * Beispiele:
 *   FirstIndexOf([]string{"a","b","c"}, "b") -> 1
 *   FirstIndexOf([]string{"x","y"}, "z")     -> -1
 */

func FirstIndexOf(list []string, needle string) int {
	if len(list) == 0 {
		return -1
	}

	for i, val := range list {
		if val == needle {
			return i
		}

	}
	return -1
}

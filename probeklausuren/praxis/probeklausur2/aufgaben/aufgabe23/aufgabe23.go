package aufgabe23

/* AUFGABENSTELLUNG (Rekursion):
 *
 * Schreiben Sie eine rekursive Funktion CountCharRuns,
 * die zählt, wie viele BLÖCKE von aufeinanderfolgenden
 * gleichen Zeichen der String enthält.
 *
 * Beispiele:
 *   "aaabbc"  -> 3   (Blöcke: "aaa", "bb", "c")
 *   "xxxx"    -> 1   (ein einziger Block)
 *   "abc"     -> 3   ("a", "b", "c")
 *   ""        -> 0
 *   "aabbbaa" -> 3   ("aa", "bbb", "aa")
 *
 * Es dürfen KEINE Schleifen verwendet werden.
 */

func CountCharRuns(s string) int {
	if s == "" {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	head := s[0]
	next := s[1]
	rest := CountCharRuns(s[1:])

	if head == next {
		return rest
	} else {

		return rest + 1
	}

}

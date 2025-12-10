package aufgabe28

/* AUFGABENSTELLUNG:
 *
 * Schreiben Sie eine Funktion CountGreaterThan, die zählt, wie viele Elemente
 * der Liste strictly größer als threshold sind.
 *
 * Beispiele:
 *   CountGreaterThan([]int{1,5,10,3}, 4) -> 2
 *   CountGreaterThan([]int{}, 10)        -> 0
 */

func CountGreaterThan(list []int, threshold int) int {
	if len(list) == 0 {
		return 0
	}
	count := 0
	for _, val := range list {

		if val > threshold {
			count++
		}
	}
	return count
}

package aufgabe22

/* AUFGABENSTELLUNG (Rekursion!):
 *
 * Schreiben Sie eine rekursive Funktion KeepEverySecondRec,
 * die aus einer int-Liste nur JEDES ZWEITE Element behält.
 *
 * Starten Sie die Zählung bei position = 1.
 *
 * Es dürfen KEINE Schleifen verwendet werden.
 *
 * Beispiele:
 *   KeepEverySecondRec([]int{1,2,3,4,5,6}, 1) -> [2 4 6]
 *   KeepEverySecondRec([]int{10,20,30}, 1)    -> [20]
 *   KeepEverySecondRec([]int{5}, 1)           -> []
 *   KeepEverySecondRec([]int{}, 1)            -> []
 */

func KeepEverySecondRec(list []int, position int) []int {
	if len(list) == 0 {
		return []int{}
	}
	head := list[0]
	rest := KeepEverySecondRec(list[1:], position+1)
	if position%2 != 0 {
		return rest
	} else {
		return append([]int{head}, rest...)
	}

}

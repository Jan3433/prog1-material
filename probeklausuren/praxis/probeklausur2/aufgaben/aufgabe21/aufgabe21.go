package aufgabe21

/* AUFGABENSTELLUNG (Rekursion):
 *
 * Schreiben Sie eine rekursive Funktion RemoveNegativesRec, die aus einer
 * int-Liste alle NEGATIVEN Zahlen entfernt.
 *
 * Es dürfen KEINE Schleifen verwendet werden.
 *
 * Beispiele:
 *   RemoveNegativesRec([]int{1,-2,3,-4,5}) -> [1 3 5]
 *   RemoveNegativesRec([]int{-1,-2,-3})     -> []
 *   RemoveNegativesRec([]int{5,4,3})        -> [5 4 3]
 *   RemoveNegativesRec([]int{})             -> []
 */

func RemoveNegativesRec(list []int) []int {

	if len(list) == 0 {
		return []int{}
	}
	head := list[0]
	rest := RemoveNegativesRec(list[1:])
	if head < 0 {
		return rest
	} else {
		return append([]int{head}, rest...)
	}

}

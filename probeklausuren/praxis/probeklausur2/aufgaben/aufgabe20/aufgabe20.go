package aufgabe20

/* AUFGABENSTELLUNG (Rekursion!):
 *
 * Schreiben Sie eine rekursive Funktion RemoveEveryThirdRec,
 * die aus einer int-Liste alle Elemente entfernt, deren
 * POSITION (1-basiert) durch 3 teilbar ist.
 *
 * Es dürfen KEINE Schleifen verwendet werden.
 *
 * Beispiele:
 *   RemoveEveryThirdRec([]int{1,2,3,4,5,6}) -> [1,2,4,5]
 *   RemoveEveryThirdRec([]int{10,20,30})    -> [10,20]
 *   RemoveEveryThirdRec([]int{})            -> []
 */

func RemoveEveryThirdRec(list []int, position int) []int {
	if len(list) == 0 {
		return []int{}
	}
	rest := RemoveEveryThirdRec(list[1:], position+1)
	head := list[0]

	if position%3 == 0 {
		return rest
	} else {

		return append([]int{head}, rest...)
	}

}

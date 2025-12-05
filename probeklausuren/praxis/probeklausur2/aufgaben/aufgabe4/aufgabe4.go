package aufgabe4

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// ElementSums erwartet zwei int-Listen l1 und l2.
// Sie liefert eine Liste, die an jeder Position
// jeweils die Summe der beiden Elemente enthält.
//
// Annahmen für die Berechnung:
// Falls eine Liste kürzer ist als die andere, soll für die Berechnung der
// hinteren Werte ihr letztes Element verwendet werden.
// Für leere Listen soll für die Berechnung ggf. 0 verwendet werden.
func ElementSums(l1, l2 []int) []int {
	result := []int{}

	current1 := 0
	current2 := 0
	max := len(l1)

	if len(l2) > max {
		max = len(l2)
	}

	for i := 0; i < max; i++ {
		// i -> max  0 1 2 3 4 5
		// l1 1,2,3,4,5 = 5
		// l2 1,2,3,4,5,6,7 = 7
		//Max = l2
		// i = 0 < 5, 7 i = 1 < 5, 7 ,.... i = 5 = 5, current1 = letzter wert +
		if len(l1) > i {
			current1 = l1[i]
		}
		if len(l2) > i {
			current2 = l2[i]
		}
		result = append(result, current1+current2)
	}

	return result
}

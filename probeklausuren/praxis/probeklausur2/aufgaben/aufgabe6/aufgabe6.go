package aufgabe6

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// DuplicateSinglets erwartet eine int-Liste list.
// Die Funktion liefert eine int-Liste, bei der alle Elemente,
// die in list nur einmal vorkommen, verdoppelt sind,
// also zwei Mal hintereinander stehen.
// Elemente, die schon in list mehrfach vorkommen, sollen wie sie sind
// ins Ergebnis übertragen werden.
func DuplicateSinglets(list []int) []int {
	dupelist := []int{}

	// äußere Schleife: über alle Elemente
	for i := 0; i < len(list); i++ {

		count := 0

		for j := 0; j < len(list); j++ {
			// TODO: Wenn list[j] gleich list[i] ist, erhöhe count
			if list[j] == list[i] {
				count++
			}
		}

		// Schritt 2: abhängig von count an dupelist anhängen
		if count == 1 {
			dupelist = append(dupelist, list[i], list[i])

		} else {
			// TODO: list[i] einmal anhängen
			dupelist = append(dupelist, list[i])
		}
	}

	return dupelist
}

package aufgabe3

import "math"

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * RANDBEDINGUNG: Die Funktion muss rekursiv sein.
 * ERREICHBARE PUNKTE: 10
 */

// CountSquares erwartet eine Liste von Zahlen.
// CountSquares liefert die Anzahl der QuadratzahlenZahlen in der Liste.

func CountSquares(list []int) int {
	if len(list) == 0 { // Basisfall
		return 0
	}

	// Prüfe das erste Element
	x := list[0]
	root := math.Sqrt(float64(x))

	if root == float64(int(root)) {
		// x ist Quadratzahl
		return 1 + CountSquares(list[1:])
	}

	// x ist KEINE Quadratzahl
	return CountSquares(list[1:])
}

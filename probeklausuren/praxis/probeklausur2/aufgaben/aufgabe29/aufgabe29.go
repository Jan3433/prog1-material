package aufgabe29

/* AUFGABENSTELLUNG:
 *
 * Gegeben ist der folgende Datentyp Entry.
 * Schreiben Sie eine Funktion UpdateEn, die den ERSTEN Eintrag in dict sucht,
 * dessen deutscher Begriff De gleich de ist, und den englischen Eintrag En
 * durch newEn ersetzt.
 *
 * Wird kein Eintrag gefunden, soll nichts verändert werden.
 */

type Entry struct {
	De string
	En string
}

func UpdateEn(dict []Entry, de string, newEn string) {

	for i := 0; i < len(dict); i++ {
		if dict[i].De == de {
			dict[i].En = newEn

		}

	}

}

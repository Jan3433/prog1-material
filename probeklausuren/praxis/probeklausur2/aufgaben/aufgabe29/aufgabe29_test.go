package aufgabe29

import "fmt"

func ExampleUpdateEn() {
	d := []Entry{
		{"Haus", "house"},
		{"Holz", "wood"},
		{"Auto", "car"},
	}

	UpdateEn(d, "Holz", "timber")
	UpdateEn(d, "Maus", "mouse")

	fmt.Println(d)

	// Output:
	// [{Haus house} {Holz timber} {Auto car}]
}

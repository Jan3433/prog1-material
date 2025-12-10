package aufgabe26

import "fmt"

func ExampleReplaceAll() {
	fmt.Println(ReplaceAll("banana", 'a', 'o'))
	fmt.Println(ReplaceAll("test", 't', 'x'))
	fmt.Println(ReplaceAll("", 'a', 'b'))

	// Output:
	// bonono
	// xesx
	//
}

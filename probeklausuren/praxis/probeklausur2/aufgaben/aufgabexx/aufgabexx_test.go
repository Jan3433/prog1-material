package aufgabeXX

import "fmt"

func ExampleRemoveChar() {
	fmt.Println(RemoveChar("banana", 'a'))
	fmt.Println(RemoveChar("hello", 'l'))
	fmt.Println(RemoveChar("aaaa", 'a'))
	fmt.Println(RemoveChar("", 'x'))

	// Output:
	// bnn
	// heo
	//
	//
}

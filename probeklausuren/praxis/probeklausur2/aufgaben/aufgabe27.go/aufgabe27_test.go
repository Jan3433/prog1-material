package aufgabe27

import "fmt"

func ExampleFirstIndexOf() {
	fmt.Println(FirstIndexOf([]string{"a", "b", "c"}, "b"))
	fmt.Println(FirstIndexOf([]string{"x", "y"}, "z"))
	fmt.Println(FirstIndexOf([]string{"apple", "banana", "pear"}, "pear"))

	// Output:
	// 1
	// -1
	// 2
}

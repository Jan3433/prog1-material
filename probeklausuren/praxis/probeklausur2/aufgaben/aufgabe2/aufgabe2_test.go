package aufgabe2

import "fmt"

func ExampleFilterDigits() {

	fmt.Println(FilterDigits("a123b"))
	fmt.Println(FilterDigits("aBcde"))
	fmt.Println(FilterDigits("0123456789"))
	fmt.Println(FilterDigits("1A9b"))
	fmt.Println(FilterDigits(""))

	// Output:
	// 123
	//
	//0123456789
	// 19
	//
}

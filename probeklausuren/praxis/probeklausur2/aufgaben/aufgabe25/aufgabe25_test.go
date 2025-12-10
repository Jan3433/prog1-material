package aufgabe25

import "fmt"

func ExampleFilterShortStrings() {
	fmt.Println(FilterShortStrings([]string{"hi", "hello", "a", "abc"}, 3))
	fmt.Println(FilterShortStrings([]string{}, 5))
	fmt.Println(FilterShortStrings([]string{"abcd", "xy"}, 2))

	// Output:
	// [hello abc]
	// []
	// [abcd xy]
}

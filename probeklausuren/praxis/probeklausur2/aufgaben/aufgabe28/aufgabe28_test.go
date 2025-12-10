package aufgabe28

import "fmt"

func ExampleCountGreaterThan() {
	fmt.Println(CountGreaterThan([]int{1, 5, 10, 3}, 4))
	fmt.Println(CountGreaterThan([]int{}, 10))
	fmt.Println(CountGreaterThan([]int{0, 1, 2, 3, 4}, 2))

	// Output:
	// 2
	// 0
	// 2
}

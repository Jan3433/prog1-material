package aufgabe21

import "fmt"

func ExampleRemoveNegativesRec() {
	fmt.Println(RemoveNegativesRec([]int{1, -2, 3, -4, 5}))
	fmt.Println(RemoveNegativesRec([]int{-1, -2, -3}))
	fmt.Println(RemoveNegativesRec([]int{5, 4, 3}))
	fmt.Println(RemoveNegativesRec([]int{}))
	fmt.Println(RemoveNegativesRec([]int{0, -1, 1, -2, 2}))

	// Output:
	// [1 3 5]
	// []
	// [5 4 3]
	// []
	// [0 1 2]
}

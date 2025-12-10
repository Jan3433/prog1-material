package aufgabe20

import "fmt"

func ExampleRemoveEveryThirdRec() {
	fmt.Println(RemoveEveryThirdRec([]int{1, 2, 3, 4, 5, 6}, 1))
	fmt.Println(RemoveEveryThirdRec([]int{10, 20, 30}, 1))
	fmt.Println(RemoveEveryThirdRec([]int{}, 1))
	fmt.Println(RemoveEveryThirdRec([]int{7, 8, 9, 10, 11}, 1))

	// Output:
	// [1 2 4 5]
	// [10 20]
	// []
	// [7 8 10 11]
}

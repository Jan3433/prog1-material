package aufgabe22

import "fmt"

func ExampleKeepEverySecondRec() {
	fmt.Println(KeepEverySecondRec([]int{1, 2, 3, 4, 5, 6}, 1))
	fmt.Println(KeepEverySecondRec([]int{10, 20, 30}, 1))
	fmt.Println(KeepEverySecondRec([]int{5}, 1))
	fmt.Println(KeepEverySecondRec([]int{}, 1))
	fmt.Println(KeepEverySecondRec([]int{7, 8, 9, 10}, 1))

	// Output:
	// [2 4 6]
	// [20]
	// []
	// []
	// [8 10]
}

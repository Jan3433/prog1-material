package aufgabe23

import "fmt"

func ExampleCountCharRuns() {
	fmt.Println(CountCharRuns("aaabbc"))
	fmt.Println(CountCharRuns("xxxx"))
	fmt.Println(CountCharRuns("abc"))
	fmt.Println(CountCharRuns(""))
	fmt.Println(CountCharRuns("aabbbaa"))

	// Output:
	// 3
	// 1
	// 3
	// 0
	// 3
	//
}

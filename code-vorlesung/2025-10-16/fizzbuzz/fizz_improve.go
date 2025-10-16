package fizzbuzz

import "fmt"

// Fizz gibt die Zahlen von 1 bis 30 aus und ersetzt dabei
// jede durch 3 teilbare Zahl durch "fizz" und jede durch
// 5 teilbare Zahl durch "buzz".
// Bei Zahlen, die sowohl durch 3 als auch durch 5 teilbar sind,
// wird "fizzbuzz" ausgegeben.
func FizzImprovedPrint(n int) {
	for i := 1; i <= n; i++ {
		if i%5 != 0 && i%3 != 0 {
			fmt.Println("i")
			continue
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		}
	}
}

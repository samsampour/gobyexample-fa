// Go از
// <a href="https://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>توابع بازگشتی</em></a>
// پشتیبانی می‌کند. این یک مثال کلاسیک است.

package main

import "fmt"

// این تابع `fact` آن‌قدر خودش را صدا می‌زند تا به حالت
// پایه، یعنی `fact(0)`، برسد.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// توابع بی‌نام هم می‌توانند بازگشتی باشند، اما برای این
	// کار باید پیش از تعریف تابع، متغیری را صریحاً با `var`
	// تعریف کنید تا تابع را در آن نگه دارد.
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		// چون `fib` قبلاً در `main` تعریف شده، Go می‌داند
		// که اینجا با `fib` کدام تابع را صدا بزند.
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}

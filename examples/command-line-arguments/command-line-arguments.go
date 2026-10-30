// [_آرگومان‌های خط فرمان_](https://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// روشی رایج برای پارامتری کردن اجرای برنامه‌ها هستند.
// برای مثال، `go run hello.go` از آرگومان‌های `run` و
// `hello.go` برای برنامهٔ `go` استفاده می‌کند.

package main

import (
	"fmt"
	"os"
)

func main() {

	// `os.Args` دسترسی به آرگومان‌های خام خط فرمان را
	// فراهم می‌کند. توجه کنید که اولین مقدار این اسلایس
	// مسیر خود برنامه است و `os.Args[1:]` آرگومان‌های
	// برنامه را نگه می‌دارد.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// آرگومان‌های جداگانه را می‌توانید با اندیس‌گذاری
	// معمولی بگیرید.
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

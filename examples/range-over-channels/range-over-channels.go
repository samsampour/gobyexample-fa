// در یک مثال [قبلی](range-over-built-in-types) دیدیم که `for` و
// `range` چطور امکان پیمایش ساختارهای داده‌ی پایه را
// فراهم می‌کنند. از همین نحو می‌توانیم برای پیمایش
// مقادیری که از یک کانال دریافت می‌شوند هم استفاده کنیم.

package main

import "fmt"

func main() {

	// روی ۲ مقدار در کانال `queue` پیمایش می‌کنیم.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// این `range` روی هر عنصر به محض دریافت از `queue`
	// پیمایش می‌کند. چون کانال را در بالا بستیم (`close`)،
	// پیمایش پس از دریافت ۲ عنصر پایان می‌یابد.
	for elem := range queue {
		fmt.Println(elem)
	}
}

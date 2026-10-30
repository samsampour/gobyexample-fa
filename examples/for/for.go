// `for` تنها ساختار حلقه در Go است. در اینجا چند نوع
// پایه‌ای از حلقه‌های `for` را می‌بینید.

package main

import "fmt"

func main() {

	// ساده‌ترین نوع، با یک شرط.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// حلقهٔ کلاسیک `for` با بخش‌های مقدار اولیه/شرط/گام.
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// روش دیگر برای تکرار ساده‌ی «این کار را N بار انجام
	// بده»، استفاده از `range` روی یک عدد صحیح است.
	for i := range 3 {
		fmt.Println("range", i)
	}

	// `for` بدون شرط آن‌قدر تکرار می‌شود تا با `break`
	// از حلقه خارج شوید یا با `return` از تابع
	// دربرگیرنده برگردید.
	for {
		fmt.Println("loop")
		break
	}

	// همچنین می‌توانید با `continue` به تکرار بعدی
	// حلقه بروید.
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

// [_توابع variadic_](https://en.wikipedia.org/wiki/Variadic_function)
// (با تعداد آرگومان متغیر) را می‌توان با هر تعدادی
// آرگومان انتهایی صدا زد. برای مثال، `fmt.Println` یک
// تابع variadic رایج است.

package main

import "fmt"

// این تابع هر تعداد دلخواهی `int` را به‌عنوان آرگومان
// می‌پذیرد.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	// داخل تابع، نوع `nums` معادل `[]int` است. می‌توانیم
	// `len(nums)` را صدا بزنیم، با `range` روی آن پیمایش
	// کنیم و غیره.
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// توابع variadic را می‌توان به روش معمول و با
	// آرگومان‌های جداگانه صدا زد.
	sum(1, 2)
	sum(1, 2, 3)

	// اگر از قبل چند آرگومان در یک اسلایس دارید، آن‌ها را
	// این‌طور با `func(slice...)` به تابع variadic بدهید.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

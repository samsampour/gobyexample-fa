// بستهٔ `slices` در Go مرتب‌سازی را برای انواع داخلی و
// انواع تعریف‌شده توسط کاربر پیاده‌سازی می‌کند. اول
// مرتب‌سازی انواع داخلی را می‌بینیم.

package main

import (
	"fmt"
	"slices"
)

func main() {

	// توابع مرتب‌سازی جنریک هستند و برای هر نوع داخلی
	// _ترتیب‌پذیر_ (ordered) کار می‌کنند. برای فهرست انواع
	// ترتیب‌پذیر، [cmp.Ordered](https://pkg.go.dev/cmp#Ordered)
	// را ببینید.
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	// نمونه‌ای از مرتب‌سازی `int`ها.
	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	// با بستهٔ `slices` می‌توانیم بررسی کنیم که آیا یک
	// اسلایس از قبل مرتب است یا نه.
	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}

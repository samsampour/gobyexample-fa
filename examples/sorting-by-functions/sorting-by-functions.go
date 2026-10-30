// گاهی می‌خواهیم یک مجموعه را بر اساس چیزی غیر از ترتیب
// طبیعی‌اش مرتب کنیم. برای مثال، فرض کنید می‌خواهیم
// رشته‌ها را به‌جای ترتیب الفبایی، بر اساس طولشان مرتب
// کنیم. این نمونه‌ای از مرتب‌سازی سفارشی در Go است.

package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	// یک تابع مقایسه برای طول رشته‌ها پیاده‌سازی می‌کنیم.
	// `cmp.Compare` برای این کار مفید است.
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	// حالا می‌توانیم `slices.SortFunc` را با این تابع مقایسهٔ
	// سفارشی صدا بزنیم تا `fruits` را بر اساس طول نام مرتب
	// کنیم.
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	// با همین تکنیک می‌توانیم اسلایسی از مقادیری را که از
	// انواع داخلی نیستند مرتب کنیم.
	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	// `people` را با `slices.SortFunc` بر اساس سن مرتب
	// می‌کنیم.
	//
	// نکته: اگر ساختار `Person` بزرگ باشد، شاید بهتر باشد
	// اسلایس شامل `*Person` باشد و تابع مرتب‌سازی را هم
	// متناسب با آن تغییر دهید. اگر مطمئن نیستید،
	// [بنچمارک](testing-and-benchmarking) بگیرید!
	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)
}

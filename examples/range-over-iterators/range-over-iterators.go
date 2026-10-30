// از نسخهٔ 1.23، Go پشتیبانی از
// [پیمایشگرها](https://go.dev/blog/range-functions)
// (iterators) را اضافه کرده که به ما اجازه می‌دهد تقریباً
// روی هر چیزی `range` بزنیم!

package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"
)

// بیایید دوباره نوع `List` از [مثال قبل](generics) را
// ببینیم. در آن مثال متد `AllElements` را داشتیم که یک
// اسلایس از همهٔ عناصر لیست برمی‌گرداند. با پیمایشگرهای
// Go می‌توانیم این کار را بهتر انجام دهیم؛ مثل کد زیر.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// All یک _پیمایشگر_ برمی‌گرداند که در Go تابعی با
// [امضای ویژه](https://pkg.go.dev/iter#Seq) است.
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		// تابع پیمایشگر تابع دیگری را به‌عنوان پارامتر
		// می‌گیرد که طبق قرارداد `yield` نامیده می‌شود (اما
		// نامش می‌تواند هر چیزی باشد). این تابع برای هر
		// عنصری که می‌خواهیم پیمایش کنیم `yield` را صدا
		// می‌زند و به مقدار بازگشتی `yield` توجه می‌کند تا در
		// صورت لزوم زودتر متوقف شود.
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// پیمایش نیازی به یک ساختار داده‌ی زیربنایی ندارد و حتی
// لازم نیست متناهی باشد! این تابع پیمایشگری روی اعداد
// فیبوناچی برمی‌گرداند: تا زمانی که `yield` مقدار `true`
// برگرداند، به کارش ادامه می‌دهد.
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// چون `List.All` یک پیمایشگر برمی‌گرداند، می‌توانیم از
	// آن در یک حلقهٔ `range` معمولی استفاده کنیم.
	for e := range lst.All() {
		fmt.Println(e)
	}

	// بسته‌هایی مثل [slices](https://pkg.go.dev/slices)
	// توابع مفید زیادی برای کار با پیمایشگرها دارند.
	// برای مثال، `Collect` هر پیمایشگری را می‌گیرد و همهٔ
	// مقادیرش را در یک اسلایس جمع می‌کند.
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	// بسته‌های کتابخانهٔ استاندارد حالا ابزارهای کمکی
	// پیمایشگر هم دارند. برای مثال، `strings.SplitSeq` روی
	// بخش‌های یک رشته پیمایش می‌کند، بدون اینکه اول یک
	// اسلایس نتیجه بسازد.
	for part := range strings.SplitSeq("go-by-example", "-") {
		fmt.Printf("part: %s\n", part)
	}

	for n := range genFib() {

		// وقتی حلقه به `break` یا یک return زودهنگام برسد،
		// تابع `yield` که به پیمایشگر داده شده `false`
		// برمی‌گرداند.
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}

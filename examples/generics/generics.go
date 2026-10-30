// از نسخهٔ 1.18، Go پشتیبانی از _جنریک‌ها_ (generics) را
// اضافه کرده است که با نام _پارامترهای نوع_
// (type parameters) هم شناخته می‌شوند.

package main

import "fmt"

// به‌عنوان نمونه‌ای از یک تابع جنریک، `SlicesIndex` یک
// اسلایس از هر نوع `comparable` و یک عنصر از همان نوع
// می‌گیرد و اندیس اولین رخداد v در s را برمی‌گرداند، یا
// اگر وجود نداشت -1 را. محدودیت `comparable` یعنی
// می‌توانیم مقادیر این نوع را با عملگرهای `==` و `!=`
// مقایسه کنیم. برای توضیح کامل‌تر این امضای نوع،
// [این مقاله](https://go.dev/blog/deconstructing-type-parameters)
// را ببینید. توجه کنید که این تابع در کتابخانهٔ
// استاندارد با نام [slices.Index](https://pkg.go.dev/slices#Index)
// وجود دارد.
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// به‌عنوان نمونه‌ای از یک نوع جنریک، `List` یک لیست
// پیوندی یک‌طرفه است که مقادیری از هر نوعی را نگه می‌دارد.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// می‌توانیم روی نوع‌های جنریک هم مثل نوع‌های معمولی متد
// تعریف کنیم، اما باید پارامترهای نوع را سر جایشان نگه
// داریم. نوع ما `List[T]` است، نه `List`.
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// AllElements همهٔ عناصر List را به‌صورت یک اسلایس
// برمی‌گرداند. در مثال بعدی روش رایج‌تری برای پیمایش همهٔ
// عناصر نوع‌های سفارشی را خواهیم دید.
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	// هنگام فراخوانی توابع جنریک، اغلب می‌توانیم به
	// _استنتاج نوع_ (type inference) تکیه کنیم. توجه کنید
	// که هنگام صدا زدن `SlicesIndex` لازم نیست نوع‌های `S`
	// و `E` را مشخص کنیم؛ کامپایلر آن‌ها را خودکار
	// تشخیص می‌دهد.
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	// ... البته می‌توانستیم آن‌ها را صریحاً هم مشخص کنیم.
	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}

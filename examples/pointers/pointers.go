// Go از <em><a href="https://en.wikipedia.org/wiki/Pointer_(computer_programming)">اشاره‌گرها</a></em>
// (pointers) پشتیبانی می‌کند که به شما اجازه می‌دهند
// ارجاع به مقادیر و رکوردها را در برنامه جابه‌جا کنید.

package main

import "fmt"

// با دو تابع `zeroval` و `zeroptr` نشان می‌دهیم که
// اشاره‌گرها در مقایسه با مقادیر چطور کار می‌کنند.
// `zeroval` یک پارامتر `int` دارد، پس آرگومان‌ها با مقدار
// (by value) به آن داده می‌شوند. `zeroval` یک کپی از
// `ival` می‌گیرد که از متغیر تابع صدازننده جداست.
func zeroval(ival int) {
	ival = 0
}

// در مقابل، `zeroptr` یک پارامتر `*int` دارد، یعنی یک
// اشاره‌گر به `int` می‌گیرد. سپس کد `*iptr` در بدنهٔ تابع،
// اشاره‌گر را _dereference_ می‌کند؛ یعنی از آدرس حافظه
// به مقدار فعلیِ آن آدرس می‌رسد. انتساب مقدار به یک
// اشاره‌گرِ dereference‌شده، مقدارِ آدرس مورد ارجاع را
// تغییر می‌دهد.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// نحو `&i` آدرس حافظهٔ `i` را می‌دهد، یعنی یک
	// اشاره‌گر به `i`.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// اشاره‌گرها را هم می‌توان چاپ کرد.
	fmt.Println("pointer:", &i)

	// با تابع داخلی `new` می‌توان یک اشاره‌گر جدید به یک
	// مقدار ساخت.
	p := new(42)
	fmt.Println("value at *p:", *p)
	zeroptr(p)
	fmt.Println("value at *p:", *p)
}

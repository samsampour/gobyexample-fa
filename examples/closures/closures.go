// Go از [_توابع بی‌نام_](https://en.wikipedia.org/wiki/Anonymous_function)
// (anonymous functions) پشتیبانی می‌کند که می‌توانند <a href="https://en.wikipedia.org/wiki/Closure_(computer_science)"><em>کلوژر</em></a> بسازند.
// توابع بی‌نام وقتی مفیدند که بخواهید تابعی را درجا
// تعریف کنید، بدون اینکه برایش نامی بگذارید.

package main

import "fmt"

// این تابع `intSeq` تابع دیگری برمی‌گرداند که آن را به
// شکل بی‌نام در بدنهٔ `intSeq` تعریف کرده‌ایم. تابع
// برگردانده‌شده متغیر `i` را _در بر می‌گیرد_
// (closes over) و به این ترتیب یک کلوژر می‌سازد.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// `intSeq` را صدا می‌زنیم و نتیجه (یک تابع) را در
	// `nextInt` می‌ریزیم. این مقدار تابعی، `i` مخصوص به
	// خودش را نگه می‌دارد که با هر بار صدا زدن `nextInt`
	// به‌روز می‌شود.
	nextInt := intSeq()

	// با چند بار صدا زدن `nextInt` اثر کلوژر را ببینید.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// برای اطمینان از اینکه این وضعیت مختص همان تابع است،
	// یک تابع جدید بسازید و امتحانش کنید.
	newInts := intSeq()
	fmt.Println(newInts())
}

// Go پشتیبانی داخلی از [عبارات منظم](https://en.wikipedia.org/wiki/Regular_expression)
// (regular expressions) دارد. این چند نمونه از کارهای
// رایج مرتبط با regexp در Go است.

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// این بررسی می‌کند که آیا یک الگو با یک رشته مطابقت دارد.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// در بالا مستقیماً از یک الگوی رشته‌ای استفاده کردیم،
	// اما برای کارهای دیگر regexp باید با `Compile` یک
	// ساختار `Regexp` بهینه‌شده بسازید.
	r, _ := regexp.Compile("p([a-z]+)ch")

	// متدهای زیادی روی این ساختارها در دسترس است. این یک
	// آزمون تطابق مثل آنچه قبلاً دیدیم است.
	fmt.Println(r.MatchString("peach"))

	// این تطابقِ regexp را پیدا می‌کند.
	fmt.Println(r.FindString("peach punch"))

	// این هم اولین تطابق را پیدا می‌کند، اما به‌جای متن
	// مطابق، اندیس‌های شروع و پایان آن را برمی‌گرداند.
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// حالت‌های `Submatch` اطلاعاتی هم دربارهٔ تطابق کل
	// الگو و هم دربارهٔ زیرتطابق‌های داخل آن می‌دهند. برای
	// مثال این کد اطلاعات هر دو `p([a-z]+)ch` و
	// `([a-z]+)` را برمی‌گرداند.
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// به همین ترتیب، این کد اطلاعات اندیس‌های تطابق‌ها و
	// زیرتطابق‌ها را برمی‌گرداند.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// حالت‌های `All` این توابع روی همهٔ تطابق‌های ورودی
	// اعمال می‌شوند، نه فقط اولی. برای مثال برای پیدا کردن
	// همهٔ تطابق‌های یک regexp.
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// این حالت‌های `All` برای توابع دیگری که بالا دیدیم هم
	// وجود دارند.
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// دادن یک عدد صحیح نامنفی به‌عنوان آرگومان دوم به این
	// توابع، تعداد تطابق‌ها را محدود می‌کند.
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// مثال‌های بالا آرگومان‌های رشته‌ای داشتند و از نام‌هایی
	// مثل `MatchString` استفاده کردند. می‌توانیم آرگومان
	// `[]byte` هم بدهیم و `String` را از نام تابع حذف کنیم.
	fmt.Println(r.Match([]byte("peach")))

	// وقتی متغیرهای سراسری با عبارات منظم می‌سازید،
	// می‌توانید از حالت `MustCompile` به‌جای `Compile`
	// استفاده کنید. `MustCompile` به‌جای برگرداندن خطا
	// panic می‌کند، که استفاده از آن را برای متغیرهای
	// سراسری امن‌تر می‌کند.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	// از بستهٔ `regexp` می‌توان برای جایگزین کردن بخش‌هایی
	// از رشته‌ها با مقادیر دیگر هم استفاده کرد.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// حالت `Func` به شما اجازه می‌دهد متن مطابق را با یک
	// تابع دلخواه تبدیل کنید.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}

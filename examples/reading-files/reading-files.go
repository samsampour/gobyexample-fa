// خواندن و نوشتن فایل‌ها از کارهای پایه‌ای است که بسیاری
// از برنامه‌های Go به آن نیاز دارند. اول چند مثال از
// خواندن فایل‌ها را می‌بینیم.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// خواندن فایل‌ها نیاز دارد که خطای بیشتر فراخوانی‌ها را
// بررسی کنیم. این تابع کمکی بررسی خطاها را در ادامه
// ساده‌تر می‌کند.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// شاید ابتدایی‌ترین کار در خواندن فایل، بارگذاری کل
	// محتوای یک فایل در حافظه باشد.
	path := filepath.Join(os.TempDir(), "dat")
	dat, err := os.ReadFile(path)
	check(err)
	fmt.Print(string(dat))

	// اغلب می‌خواهید کنترل بیشتری روی نحوهٔ خواندن و اینکه
	// کدام بخش‌های فایل خوانده شوند داشته باشید. برای این
	// کارها، اول فایل را با `Open` باز کنید تا یک مقدار
	// `os.File` بگیرید.
	f, err := os.Open(path)
	check(err)

	// چند بایت از ابتدای فایل می‌خوانیم. اجازهٔ خواندن تا
	// ۵ بایت را می‌دهیم، اما به تعداد بایت‌هایی که واقعاً
	// خوانده شد هم توجه می‌کنیم.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// همچنین می‌توانید با `Seek` به موقعیت مشخصی در فایل
	// بروید و از آنجا `Read` کنید.
	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// روش‌های دیگر جابه‌جایی (seek) نسبت به موقعیت فعلی
	// مکان‌نما هستند،
	_, err = f.Seek(2, io.SeekCurrent)
	check(err)

	// و نسبت به انتهای فایل.
	_, err = f.Seek(-4, io.SeekEnd)
	check(err)

	// بستهٔ `io` توابعی دارد که ممکن است برای خواندن فایل
	// مفید باشند. برای مثال، خواندن‌هایی مثل موارد بالا را
	// می‌توان با `ReadAtLeast` مطمئن‌تر پیاده‌سازی کرد.
	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// تابع داخلی برای برگشت به ابتدا (rewind) وجود ندارد،
	// اما `Seek(0, io.SeekStart)` همین کار را انجام می‌دهد.
	_, err = f.Seek(0, io.SeekStart)
	check(err)

	// بستهٔ `bufio` یک خوانندهٔ بافردار پیاده‌سازی می‌کند که
	// هم به دلیل کارایی‌اش در خواندن‌های کوچکِ زیاد و هم به
	// دلیل متدهای خواندن اضافه‌ای که دارد می‌تواند مفید باشد.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// وقتی کارتان تمام شد فایل را ببندید (معمولاً این کار
	// بلافاصله پس از `Open` با `defer` زمان‌بندی می‌شود).
	f.Close()
}

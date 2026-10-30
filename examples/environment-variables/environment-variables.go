// [متغیرهای محیطی](https://en.wikipedia.org/wiki/Environment_variable)
// سازوکاری همگانی برای
// [انتقال اطلاعات پیکربندی به برنامه‌های یونیکس](https://www.12factor.net/config)
// هستند. ببینیم چطور متغیرهای محیطی را تنظیم کنیم، بخوانیم
// و فهرست کنیم.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// برای تنظیم یک جفت کلید/مقدار از `os.Setenv` استفاده
	// کنید. برای خواندن مقدار یک کلید از `os.Getenv`
	// استفاده کنید. اگر کلید در محیط وجود نداشته باشد،
	// رشتهٔ خالی برمی‌گرداند.
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// با `os.Environ` همهٔ جفت‌های کلید/مقدار محیط را
	// فهرست کنید. این تابع اسلایسی از رشته‌ها به شکل
	// `KEY=value` برمی‌گرداند. می‌توانید آن‌ها را با
	// `strings.SplitN` جدا کنید تا کلید و مقدار را بگیرید.
	// اینجا همهٔ کلیدها را چاپ می‌کنیم.
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}

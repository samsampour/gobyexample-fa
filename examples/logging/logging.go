// کتابخانهٔ استاندارد Go ابزارهای سرراستی برای تولید لاگ
// در برنامه‌های Go دارد: بستهٔ [log](https://pkg.go.dev/log)
// برای خروجی آزاد و بستهٔ
// [log/slog](https://pkg.go.dev/log/slog) برای خروجی
// ساخت‌یافته (structured).
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {

	// صدا زدن ساده‌ی توابعی مثل `Println` از بستهٔ `log`
	// از لاگرِ _استاندارد_ استفاده می‌کند که از قبل برای
	// خروجی لاگ معقول در `os.Stderr` پیکربندی شده است.
	// متدهای دیگری مثل `Fatal*` یا `Panic*` پس از ثبت لاگ
	// از برنامه خارج می‌شوند.
	log.Println("standard logger")

	// لاگرها را می‌توان با _فلگ‌ها_ پیکربندی کرد تا قالب
	// خروجی‌شان تعیین شود. به‌طور پیش‌فرض، لاگر استاندارد
	// فلگ‌های `log.Ldate` و `log.Ltime` را دارد که در
	// `log.LstdFlags` جمع شده‌اند. برای مثال می‌توانیم
	// فلگ‌هایش را تغییر دهیم تا زمان را با دقت میکروثانیه
	// ثبت کند.
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// همچنین از ثبت نام فایل و شماره خطی که تابع `log` از
	// آن صدا زده شده پشتیبانی می‌کند.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// ممکن است مفید باشد یک لاگر سفارشی بسازیم و آن را
	// جابه‌جا کنیم. هنگام ساختن لاگر جدید می‌توانیم یک
	// _پیشوند_ (prefix) تعیین کنیم تا خروجی‌اش از لاگرهای
	// دیگر قابل تشخیص باشد.
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	// با متد `SetPrefix` می‌توانیم پیشوند لاگرهای موجود
	// (از جمله لاگر استاندارد) را هم تنظیم کنیم.
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// لاگرها می‌توانند مقصدهای خروجی سفارشی داشته باشند؛
	// هر `io.Writer`ی کار می‌کند.
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	// این فراخوانی خروجی لاگ را در `buf` می‌نویسد.
	buflog.Println("hello")

	// این یکی آن را واقعاً در خروجی استاندارد نشان می‌دهد.
	fmt.Print("from buflog:", buf.String())

	// بستهٔ `slog` خروجی لاگِ _ساخت‌یافته_ فراهم می‌کند.
	// برای مثال، ثبت لاگ با قالب JSON ساده است.
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	// علاوه بر پیام، خروجی `slog` می‌تواند هر تعداد دلخواه
	// جفت key=value داشته باشد.
	myslog.Info("hello again", "key", "val", "age", 25)
}

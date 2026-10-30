// در مثال قبل راه‌اندازی یک [سرور HTTP](http-server) ساده
// را دیدیم. سرورهای HTTP برای نشان دادن کاربرد
// `context.Context` در کنترل لغو (cancellation) مفیدند.
// یک `Context` مهلت‌های زمانی (deadlines)، سیگنال‌های لغو
// و دیگر مقادیر مربوط به یک درخواست را از مرزهای API و
// بین گوروتین‌ها منتقل می‌کند.
package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// سازوکار `net/http` برای هر درخواست یک
	// `context.Context` می‌سازد که با متد `Context()` در
	// دسترس است.
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// چند ثانیه پیش از فرستادن پاسخ به کلاینت صبر می‌کنیم.
	// این می‌تواند کاری را که سرور انجام می‌دهد شبیه‌سازی
	// کند. در حین کار، حواسمان به کانال `Done()` در context
	// هست تا اگر سیگنالی رسید که باید کار را لغو کنیم، هر
	// چه زودتر برگردیم.
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// متد `Err()` در context خطایی برمی‌گرداند که توضیح
		// می‌دهد چرا کانال `Done()` بسته شده است.
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	// مثل قبل، هندلرمان را روی مسیر "/hello" ثبت می‌کنیم
	// و سرویس‌دهی را شروع می‌کنیم.
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}

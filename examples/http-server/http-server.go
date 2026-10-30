// نوشتن یک سرور HTTP ساده با بستهٔ `net/http` آسان است.
package main

import (
	"fmt"
	"net/http"
)

// یکی از مفاهیم بنیادی در سرورهای `net/http`،
// *هندلرها* (handlers) هستند. هندلر شیئی است که اینترفیس
// `http.Handler` را پیاده‌سازی می‌کند. یک روش رایج برای
// نوشتن هندلر، استفاده از تبدیل‌گر `http.HandlerFunc` روی
// توابعی با امضای مناسب است.
func hello(w http.ResponseWriter, req *http.Request) {

	// توابعی که نقش هندلر دارند، یک `http.ResponseWriter` و
	// یک `http.Request` به‌عنوان آرگومان می‌گیرند. از
	// response writer برای پر کردن پاسخ HTTP استفاده
	// می‌شود. اینجا پاسخ ساده‌ی ما فقط "hello\n" است.
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	// این هندلر کار کمی پیچیده‌تری انجام می‌دهد: همهٔ
	// سرآیندهای (headers) درخواست HTTP را می‌خواند و آن‌ها
	// را در بدنهٔ پاسخ برمی‌گرداند.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// هندلرهایمان را با تابع کمکی `http.HandleFunc` روی
	// مسیرهای سرور ثبت می‌کنیم. این تابع *مسیریاب پیش‌فرض*
	// (default router) بستهٔ `net/http` را تنظیم می‌کند و یک
	// تابع را به‌عنوان آرگومان می‌گیرد.
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// در نهایت، `ListenAndServe` را با پورت و یک هندلر صدا
	// می‌زنیم. `nil` به آن می‌گوید از همان مسیریاب پیش‌فرضی
	// که الان تنظیم کردیم استفاده کند.
	http.ListenAndServe(":8090", nil)
}

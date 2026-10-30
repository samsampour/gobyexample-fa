// وقتی از کانال‌ها به‌عنوان پارامتر تابع استفاده می‌کنید،
// می‌توانید مشخص کنید که کانال فقط برای ارسال است یا فقط
// برای دریافت. این مشخص بودن، ایمنی نوع برنامه را
// افزایش می‌دهد.

package main

import "fmt"

// این تابع `ping` فقط کانالی برای ارسال مقادیر می‌پذیرد.
// تلاش برای دریافت از این کانال، خطای زمان کامپایل
// خواهد بود.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// تابع `pong` یک کانال برای دریافت (`pings`) و یک کانال
// دوم برای ارسال (`pongs`) می‌پذیرد.
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

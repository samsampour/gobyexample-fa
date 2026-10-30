// [_محدودسازی نرخ_](https://en.wikipedia.org/wiki/Rate_limiting)
// (rate limiting) سازوکار مهمی برای کنترل مصرف منابع و
// حفظ کیفیت سرویس است. Go با گوروتین‌ها، کانال‌ها و
// [تیکرها](tickers) به شکلی زیبا از محدودسازی نرخ
// پشتیبانی می‌کند.

package main

import (
	"fmt"
	"time"
)

func main() {

	// اول محدودسازی نرخ ساده را می‌بینیم. فرض کنید
	// می‌خواهیم رسیدگی به درخواست‌های ورودی را محدود کنیم.
	// این درخواست‌ها را از کانالی به همین نام سرویس
	// می‌دهیم.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// کانال `limiter` هر ۲۰۰ میلی‌ثانیه یک مقدار دریافت
	// می‌کند. این تنظیم‌کنندهٔ طرح محدودسازی نرخ ماست.
	limiter := time.Tick(200 * time.Millisecond)

	// با مسدود شدن روی دریافت از کانال `limiter` پیش از
	// سرویس دادن به هر درخواست، خودمان را به ۱ درخواست در
	// هر ۲۰۰ میلی‌ثانیه محدود می‌کنیم.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// ممکن است بخواهیم در طرح محدودسازی نرخ، اجازهٔ
	// رگبارهای کوتاهی از درخواست‌ها (burst) را بدهیم و در
	// عین حال محدودیت نرخ کلی حفظ شود. این کار را می‌توان
	// با بافردار کردن کانال limiter انجام داد. کانال
	// `burstyLimiter` اجازهٔ رگبارهایی تا ۳ رویداد را می‌دهد.
	burstyLimiter := make(chan time.Time, 3)

	// کانال را پر می‌کنیم تا رگبار مجاز را نشان دهد.
	for range 3 {
		burstyLimiter <- time.Now()
	}

	// هر ۲۰۰ میلی‌ثانیه سعی می‌کنیم یک مقدار جدید به
	// `burstyLimiter` اضافه کنیم، تا سقف ۳.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// حالا ۵ درخواست ورودی دیگر را شبیه‌سازی می‌کنیم. ۳
	// درخواست اول از قابلیت رگبارِ `burstyLimiter` بهره
	// می‌برند.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

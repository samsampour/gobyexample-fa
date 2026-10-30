// Go پشتیبانی گسترده‌ای از زمان‌ها و بازه‌های زمانی
// (durations) دارد؛ این چند نمونه است.

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// با گرفتن زمان فعلی شروع می‌کنیم.
	now := time.Now()
	p(now)

	// می‌توانید با دادن سال، ماه، روز و غیره یک ساختار
	// `time` بسازید. زمان‌ها همیشه به یک `Location`، یعنی
	// منطقهٔ زمانی، وابسته‌اند.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// همان‌طور که انتظار دارید می‌توانید اجزای مختلف مقدار
	// زمان را استخراج کنید.
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// روز هفته (`Weekday`) از دوشنبه تا یکشنبه هم در
	// دسترس است.
	p(then.Weekday())

	// این متدها دو زمان را مقایسه می‌کنند و به‌ترتیب بررسی
	// می‌کنند که آیا اولی قبل، بعد، یا هم‌زمان با دومی است.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// متد `Sub` یک `Duration` برمی‌گرداند که فاصلهٔ بین
	// دو زمان را نشان می‌دهد.
	diff := now.Sub(then)
	p(diff)

	// می‌توانیم طول این بازه را در واحدهای مختلف حساب کنیم.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// با `Add` می‌توانید زمان را به اندازهٔ یک بازهٔ مشخص
	// جلو ببرید، یا با `-` به اندازهٔ یک بازه عقب ببرید.
	p(then.Add(diff))
	p(then.Add(-diff))
}

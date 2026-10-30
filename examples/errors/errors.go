// در Go روش رایج این است که خطاها را از طریق یک مقدار
// بازگشتی صریح و جداگانه اعلام کنیم. این برخلاف
// استثناهایی (exceptions) است که در زبان‌هایی مثل Java،
// Python و Ruby استفاده می‌شود، و برخلاف مقدار واحدِ
// نتیجه/خطا که گاهی در C به کار می‌رود. رویکرد Go باعث
// می‌شود به‌راحتی ببینیم کدام توابع خطا برمی‌گردانند و
// خطاها را با همان ساختارهای زبانی مدیریت کنیم که برای
// کارهای دیگر (غیر از خطا) استفاده می‌کنیم.
//
// برای جزئیات بیشتر، مستندات [بستهٔ errors](https://pkg.go.dev/errors)
// و [این مقاله](https://go.dev/blog/go1.13-errors) را
// ببینید.

package main

import (
	"errors"
	"fmt"
)

// طبق قرارداد، خطاها آخرین مقدار بازگشتی هستند و نوع
// `error` دارند که یک اینترفیس داخلی است.
func f(arg int) (int, error) {
	if arg == 42 {
		// `errors.New` یک مقدار `error` ساده با پیام خطای
		// داده‌شده می‌سازد.
		return -1, errors.New("can't work with 42")
	}

	// مقدار `nil` در جایگاه خطا یعنی هیچ خطایی رخ
	// نداده است.
	return arg + 3, nil
}

// خطای نگهبان (sentinel error) متغیری از پیش تعریف‌شده است
// که برای نشان دادن یک وضعیت خطای مشخص به کار می‌رود.
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {

		// می‌توانیم خطاها را در خطاهای سطح بالاتر بپیچیم
		// (wrap) تا اطلاعات زمینه‌ای اضافه کنیم. ساده‌ترین
		// راه، استفاده از فعل `%w` در `fmt.Errorf` است.
		// خطاهای پیچیده‌شده یک زنجیرهٔ منطقی می‌سازند (A، B را
		// می‌پیچد که C را می‌پیچد و ...) که می‌توان با توابعی
		// مثل `errors.Is` و `errors.AsType` آن را بررسی کرد.
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {

		// روش رایج، بررسی درون‌خطی خطا در همان خط `if`
		// است.
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {

			// `errors.Is` بررسی می‌کند که آیا یک خطای مشخص (یا هر خطایی
			// در زنجیره‌اش) با یک مقدار خطای مشخص مطابقت دارد. این
			// مخصوصاً برای خطاهای پیچیده یا تودرتو مفید است و به شما
			// اجازه می‌دهد نوع خاصی از خطا یا خطاهای نگهبان را در
			// زنجیرهٔ خطاها پیدا کنید.
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}
}

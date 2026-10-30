// می‌توان با پیاده‌سازی متد `Error()` روی یک نوع، نوع خطای
// سفارشی تعریف کرد. این نسخه‌ای از مثال قبل است که از
// یک نوع سفارشی برای نمایش صریح خطای آرگومان استفاده
// می‌کند.

package main

import (
	"errors"
	"fmt"
)

// نام نوع خطای سفارشی معمولاً با پسوند «Error» تمام می‌شود.
type argError struct {
	arg     int
	message string
}

// افزودن این متد `Error` باعث می‌شود `argError` اینترفیس
// `error` را پیاده‌سازی کند.
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {

		// خطای سفارشی خودمان را برمی‌گردانیم.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// `errors.AsType` نسخهٔ پیشرفته‌تر `errors.Is` است.
	// بررسی می‌کند که آیا یک خطای مشخص (یا هر خطایی در
	// زنجیره‌اش) با یک نوع خطای مشخص مطابقت دارد، و در
	// این صورت آن را به مقداری از همان نوع تبدیل می‌کند و
	// `true` هم برمی‌گرداند. اگر مطابقتی نباشد، مقدار
	// بازگشتی دوم `false` است.
	_, err := f(42)
	if ae, ok := errors.AsType[*argError](err); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}

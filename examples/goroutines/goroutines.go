// _گوروتین_ (goroutine) یک رشتهٔ اجرایی (thread) سبک است.

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// فرض کنید فراخوانی تابع `f(s)` را داریم. این روش
	// معمولِ صدا زدن آن است که به‌صورت همگام (synchronous)
	// اجرا می‌شود.
	f("direct")

	// برای اجرای این تابع در یک گوروتین، از `go f(s)`
	// استفاده کنید. این گوروتین جدید هم‌زمان با گوروتینِ
	// صدازننده اجرا می‌شود.
	go f("goroutine")

	// می‌توانید برای فراخوانی یک تابع بی‌نام هم گوروتین
	// راه بیندازید.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// حالا دو فراخوانی تابع ما به‌صورت ناهمگام در
	// گوروتین‌های جداگانه اجرا می‌شوند. منتظر می‌مانیم تا
	// تمام شوند (برای روشی مطمئن‌تر، از
	// [WaitGroup](waitgroups) استفاده کنید).
	time.Sleep(time.Second)
	fmt.Println("done")
}

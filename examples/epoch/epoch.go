// یکی از نیازهای رایج در برنامه‌ها، گرفتن تعداد ثانیه‌ها،
// میلی‌ثانیه‌ها یا نانوثانیه‌های گذشته از
// [مبدأ زمانی یونیکس](https://en.wikipedia.org/wiki/Unix_time)
// (Unix epoch) است. این روش انجامش در Go است.

package main

import (
	"fmt"
	"time"
)

func main() {

	// از `time.Now` همراه با `Unix`، `UnixMilli` یا
	// `UnixNano` استفاده کنید تا زمان گذشته از مبدأ یونیکس
	// را به‌ترتیب به ثانیه، میلی‌ثانیه یا نانوثانیه بگیرید.
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// می‌توانید تعداد صحیح ثانیه‌ها یا نانوثانیه‌های گذشته
	// از مبدأ را هم به `time` متناظر تبدیل کنید.
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}

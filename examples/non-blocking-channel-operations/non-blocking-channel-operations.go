// ارسال و دریافت ساده روی کانال‌ها مسدودکننده است. اما
// می‌توانیم از `select` همراه با بخش `default` استفاده
// کنیم تا ارسال و دریافت _غیرمسدودکننده_ و حتی `select`
// چندراهه‌ی غیرمسدودکننده پیاده‌سازی کنیم.

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// این یک دریافت غیرمسدودکننده است. اگر مقداری روی
	// `messages` در دسترس باشد، `select` حالت
	// `<-messages` را با آن مقدار انتخاب می‌کند. در غیر
	// این صورت بلافاصله حالت `default` را انتخاب می‌کند.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// ارسال غیرمسدودکننده هم به همین شکل کار می‌کند.
	// اینجا `msg` نمی‌تواند به کانال `messages` فرستاده
	// شود، چون کانال بافر ندارد و گیرنده‌ای هم وجود ندارد.
	// بنابراین حالت `default` انتخاب می‌شود.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// می‌توانیم چند `case` را بالای بخش `default` بگذاریم
	// تا یک select چندراهه‌ی غیرمسدودکننده بسازیم. اینجا
	// سعی می‌کنیم به‌صورت غیرمسدودکننده هم از `messages` و
	// هم از `signals` دریافت کنیم.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

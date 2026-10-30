// دستور _select_ در Go به شما اجازه می‌دهد منتظر چند
// عملیات کانال به‌طور هم‌زمان بمانید. ترکیب گوروتین‌ها
// و کانال‌ها با select یکی از قابلیت‌های قدرتمند Go است.

package main

import (
	"fmt"
	"time"
)

func main() {

	// در این مثال روی دو کانال select می‌کنیم.
	c1 := make(chan string)
	c2 := make(chan string)

	// هر کانال پس از مدتی یک مقدار دریافت می‌کند تا مثلاً
	// عملیات RPC مسدودکننده‌ای را شبیه‌سازی کنیم که در
	// گوروتین‌های هم‌روند اجرا می‌شوند.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// از `select` استفاده می‌کنیم تا هم‌زمان منتظر هر دو
	// مقدار بمانیم و هر کدام را به محض رسیدن چاپ کنیم.
	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

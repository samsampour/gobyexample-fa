// در مثال قبل دیدیم که چطور وضعیت یک شمارندهٔ ساده را با
// [عملیات اتمیک](atomic-counters) مدیریت کنیم. برای
// وضعیت‌های پیچیده‌تر می‌توانیم از یک [_میوتکس_](https://en.wikipedia.org/wiki/Mutual_exclusion)
// (mutex) استفاده کنیم تا دسترسی به داده‌ها از چند
// گوروتین امن باشد.

package main

import (
	"fmt"
	"sync"
)

// Container یک مپ از شمارنده‌ها نگه می‌دارد؛ چون می‌خواهیم
// آن را هم‌زمان از چند گوروتین به‌روز کنیم، یک `Mutex`
// اضافه می‌کنیم تا دسترسی را همگام کند.
// توجه کنید که میوتکس‌ها نباید کپی شوند، پس اگر این
// `struct` جابه‌جا می‌شود، باید با اشاره‌گر باشد.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// پیش از دسترسی به `counters` میوتکس را قفل می‌کنیم و
	// در پایان تابع با دستور [defer](defer) قفل را باز
	// می‌کنیم.
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		// توجه کنید که مقدار صفر میوتکس همان‌طور که هست قابل
		// استفاده است، پس اینجا نیازی به مقداردهی اولیه نیست.
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// این تابع یک شمارندهٔ نام‌دار را در یک حلقه افزایش
	// می‌دهد.
	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
	}

	// چند گوروتین را هم‌زمان اجرا می‌کنیم؛ توجه کنید که
	// همهٔ آن‌ها به یک `Container` دسترسی دارند و دو تا از
	// آن‌ها به یک شمارندهٔ مشترک.
	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("b", 10000)
	})

	// صبر می‌کنیم تا گوروتین‌ها تمام شوند
	wg.Wait()
	fmt.Println(c.counters)
}

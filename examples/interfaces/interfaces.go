// _اینترفیس‌ها_ (interfaces) مجموعه‌هایی نام‌دار از امضای
// متدها هستند.

package main

import (
	"fmt"
	"math"
)

// این یک اینترفیس ساده برای اشکال هندسی است.
type geometry interface {
	area() float64
	perim() float64
}

// در این مثال، این اینترفیس را روی نوع‌های `rect` و
// `circle` پیاده‌سازی می‌کنیم.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// برای پیاده‌سازی یک اینترفیس در Go کافی است همهٔ متدهای
// آن اینترفیس را پیاده‌سازی کنیم. اینجا `geometry` را
// روی `rect` پیاده‌سازی می‌کنیم.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// پیاده‌سازی برای `circle`.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// اگر یک متغیر از نوع اینترفیس باشد، می‌توانیم متدهای آن
// اینترفیس را رویش صدا بزنیم. این تابع عمومی `measure`
// از همین ویژگی استفاده می‌کند تا روی هر `geometry`ای کار
// کند.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// گاهی دانستن نوع زمان اجرای یک مقدار اینترفیسی مفید است.
// یک راه، استفاده از *type assertion* است که اینجا
// می‌بینید؛ راه دیگر یک [`switch` نوعی](switch) است.
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// هر دو نوع ساختار `circle` و `rect` اینترفیس
	// `geometry` را پیاده‌سازی می‌کنند، پس می‌توانیم
	// نمونه‌هایی از
	// این ساختارها را به‌عنوان آرگومان به `measure` بدهیم.
	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}

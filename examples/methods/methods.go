// Go از _متدهایی_ پشتیبانی می‌کند که روی انواع ساختار
// تعریف می‌شوند.

package main

import "fmt"

type rect struct {
	width, height int
}

// این متد `area` یک _نوع گیرنده_ (receiver type) از جنس
// `*rect` دارد.
func (r *rect) area() int {
	return r.width * r.height
}

// متدها را می‌توان هم برای گیرندهٔ اشاره‌گری و هم برای
// گیرندهٔ مقداری تعریف کرد. این یک نمونه از گیرندهٔ
// مقداری است.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// اینجا دو متدی را که برای ساختارمان تعریف کردیم صدا
	// می‌زنیم.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go در فراخوانی متدها، تبدیل بین مقدار و اشاره‌گر را
	// خودکار انجام می‌دهد. ممکن است بخواهید از گیرندهٔ
	// اشاره‌گری استفاده کنید تا هنگام فراخوانی متد کپی
	// انجام نشود، یا متد بتواند ساختار گیرنده را تغییر دهد.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}

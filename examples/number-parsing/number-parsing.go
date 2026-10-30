// تجزیهٔ اعداد از رشته‌ها کاری ساده اما رایج در بسیاری از
// برنامه‌هاست؛ این روش انجامش در Go است.

package main

// بستهٔ داخلی `strconv` امکان تجزیهٔ اعداد را فراهم
// می‌کند.
import (
	"fmt"
	"strconv"
)

func main() {

	// در `ParseFloat`، این `64` تعیین می‌کند که با چند بیت
	// دقت تجزیه شود.
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// در `ParseInt`، عدد `0` یعنی مبنا از روی خود رشته
	// تشخیص داده شود. `64` یعنی نتیجه باید در ۶۴ بیت جا
	// شود.
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// `ParseInt` اعداد با قالب هگز را هم تشخیص می‌دهد.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// `ParseUint` هم در دسترس است.
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// `Atoi` تابعی کمکی برای تجزیهٔ ساده‌ی `int` در مبنای
	// ۱۰ است.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// توابع تجزیه برای ورودی نادرست خطا برمی‌گردانند.
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}

// URLها [روشی یکسان برای پیدا کردن منابع](https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/)
// فراهم می‌کنند. این روش تجزیهٔ URL در Go است.

package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// این URL نمونه را تجزیه می‌کنیم که شامل scheme،
	// اطلاعات احراز هویت، میزبان (host)، پورت، مسیر،
	// پارامترهای کوئری و fragment است.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// URL را تجزیه می‌کنیم و مطمئن می‌شویم خطایی نیست.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// دسترسی به scheme ساده است.
	fmt.Println(u.Scheme)

	// `User` همهٔ اطلاعات احراز هویت را دارد؛ برای گرفتن
	// مقادیر جداگانه، `Username` و `Password` را روی آن صدا
	// بزنید.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// `Host` هم نام میزبان و هم پورت (در صورت وجود) را در
	// بر دارد. برای جدا کردنشان از `SplitHostPort` استفاده
	// کنید.
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// اینجا `path` و fragment بعد از `#` را استخراج
	// می‌کنیم.
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// برای گرفتن پارامترهای کوئری به شکل رشتهٔ `k=v`، از
	// `RawQuery` استفاده کنید. همچنین می‌توانید پارامترهای
	// کوئری را به یک مپ تجزیه کنید. مپ‌های حاصل از رشته به
	// اسلایسی از رشته‌ها هستند، پس اگر فقط اولین مقدار را
	// می‌خواهید، از اندیس `[0]` استفاده کنید.
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}

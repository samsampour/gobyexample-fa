// بستهٔ `strings` در کتابخانهٔ استاندارد توابع مفید زیادی
// برای کار با رشته‌ها دارد. این چند مثال برای آشنایی با
// این بسته است.

package main

import (
	"fmt"
	s "strings"
)

// برای `fmt.Println` نام کوتاه‌تری می‌گذاریم، چون در ادامه
// زیاد از آن استفاده می‌کنیم.
var p = fmt.Println

func main() {

	// این نمونه‌ای از توابع موجود در `strings` است. چون
	// این‌ها توابعِ بسته هستند، نه متدهای خود شیء رشته،
	// باید رشتهٔ مورد نظر را به‌عنوان آرگومان اول به تابع
	// بدهیم. توابع بیشتری را می‌توانید در مستندات بستهٔ
	// [`strings`](https://pkg.go.dev/strings) پیدا کنید.
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}

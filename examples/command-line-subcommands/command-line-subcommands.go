// برخی ابزارهای خط فرمان، مثل ابزار `go` یا `git`،
// *زیرفرمان‌های* (subcommands) زیادی دارند که هر کدام
// مجموعه فلگ‌های خودش را دارد. برای مثال، `go build` و
// `go get` دو زیرفرمان متفاوت ابزار `go` هستند.
// بستهٔ `flag` به ما اجازه می‌دهد به‌راحتی زیرفرمان‌های
// ساده‌ای تعریف کنیم که فلگ‌های خودشان را دارند.

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// یک زیرفرمان را با تابع `NewFlagSet` تعریف می‌کنیم و
	// سپس فلگ‌های جدیدِ مخصوص این زیرفرمان را تعریف
	// می‌کنیم.
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// برای یک زیرفرمان دیگر می‌توانیم فلگ‌های متفاوتی
	// تعریف کنیم.
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// انتظار می‌رود زیرفرمان اولین آرگومان برنامه باشد.
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// بررسی می‌کنیم کدام زیرفرمان صدا زده شده.
	switch os.Args[1] {

	// برای هر زیرفرمان، فلگ‌های خودش را تجزیه می‌کنیم و به
	// آرگومان‌های موقعیتی انتهایی دسترسی داریم.
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}

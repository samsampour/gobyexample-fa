// گاهی می‌خواهیم برنامه‌های Go ما
// [سیگنال‌های یونیکس](https://en.wikipedia.org/wiki/Unix_signal)
// را هوشمندانه مدیریت کنند. برای مثال، ممکن است بخواهیم
// یک سرور با دریافت `SIGTERM` به‌آرامی (gracefully) خاموش
// شود، یا یک ابزار خط فرمان با دریافت `SIGINT` پردازش
// ورودی را متوقف کند. این یک روش مدرن برای مدیریت
// سیگنال‌ها با context است.

package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
)

func main() {
	// `signal.NotifyContext` یک context برمی‌گرداند که با
	// رسیدن یکی از سیگنال‌های فهرست‌شده لغو می‌شود.
	ctx, stop := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// برنامه اینجا منتظر می‌ماند تا یکی از سیگنال‌های
	// تنظیم‌شده دریافت شود.
	fmt.Println("awaiting signal")
	<-ctx.Done()

	// `context.Cause` گزارش می‌دهد که چرا context لغو شده.
	// برای لغوی که با سیگنال رخ داده، شامل مقدار سیگنال
	// هم هست.
	fmt.Println()
	fmt.Println(context.Cause(ctx))
	fmt.Println("exiting")
}

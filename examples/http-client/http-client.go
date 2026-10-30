// کتابخانهٔ استاندارد Go در بستهٔ `net/http` پشتیبانی
// بسیار خوبی از کلاینت‌ها و سرورهای HTTP دارد. در این
// مثال از آن برای ارسال درخواست‌های ساده‌ی HTTP استفاده
// می‌کنیم.
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// یک درخواست HTTP GET به یک سرور می‌فرستیم. `http.Get`
	// میان‌بری کاربردی به‌جای ساختن شیء `http.Client` و صدا
	// زدن متد `Get` آن است؛ از شیء `http.DefaultClient`
	// استفاده می‌کند که تنظیمات پیش‌فرض مفیدی دارد.
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// وضعیت پاسخ HTTP را چاپ می‌کنیم.
	fmt.Println("Response status:", resp.Status)

	// ۵ خط اول بدنهٔ پاسخ را چاپ می‌کنیم.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

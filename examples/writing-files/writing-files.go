// نوشتن فایل در Go از الگوهایی شبیه آنچه قبلاً برای
// خواندن دیدیم پیروی می‌کند.

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// برای شروع، این روش ریختن یک رشته (یا فقط بایت‌ها) در
	// یک فایل است.
	d1 := []byte("hello\ngo\n")
	path1 := filepath.Join(os.TempDir(), "dat1")
	err := os.WriteFile(path1, d1, 0644)
	check(err)

	// برای نوشتن‌های جزئی‌تر، یک فایل را برای نوشتن باز
	// کنید.
	path2 := filepath.Join(os.TempDir(), "dat2")
	f, err := os.Create(path2)
	check(err)

	// روش رایج این است که بلافاصله پس از باز کردن فایل،
	// `Close` را defer کنیم.
	defer f.Close()

	// همان‌طور که انتظار دارید می‌توانید اسلایس‌های بایت را
	// با `Write` بنویسید.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// `WriteString` هم در دسترس است.
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// با `Sync` نوشته‌ها را روی حافظهٔ پایدار (دیسک) ثبت
	// کنید.
	f.Sync()

	// `bufio` علاوه بر خواننده‌های بافرداری که قبلاً دیدیم،
	// نویسنده‌های بافردار هم فراهم می‌کند.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// با `Flush` مطمئن شوید که همهٔ عملیات بافرشده روی
	// نویسندهٔ زیربنایی اعمال شده‌اند.
	w.Flush()

}

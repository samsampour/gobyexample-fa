// از _defer_ استفاده می‌شود تا مطمئن شویم یک فراخوانی
// تابع در مراحل بعدی اجرای برنامه انجام می‌شود؛ معمولاً
// برای پاک‌سازی (cleanup). `defer` اغلب جایی به کار
// می‌رود که در زبان‌های دیگر مثلاً از `ensure` و `finally`
// استفاده می‌شود.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// فرض کنید می‌خواهیم یک فایل بسازیم، در آن بنویسیم و در
// پایان کار آن را ببندیم. این کار را می‌توانیم این‌طور
// با `defer` انجام دهیم.
func main() {

	// بلافاصله پس از گرفتن شیء فایل با `createFile`، بستن
	// آن فایل را با `closeFile` به تعویق می‌اندازیم (defer).
	// این فراخوانی در پایان تابع دربرگیرنده (`main`) و پس
	// از تمام شدن `writeFile` اجرا می‌شود.
	path := filepath.Join(os.TempDir(), "defer.txt")
	f := createFile(path)
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	// بررسی خطا هنگام بستن فایل مهم است، حتی در یک تابع
	// defer‌شده.
	if err != nil {
		panic(err)
	}
}

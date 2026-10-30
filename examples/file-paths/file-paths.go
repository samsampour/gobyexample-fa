// بستهٔ `filepath` توابعی برای تجزیه و ساختن *مسیرهای
// فایل* فراهم می‌کند، به شکلی که بین سیستم‌عامل‌های مختلف
// قابل‌حمل باشد؛ برای مثال `dir/file` در لینوکس در برابر
// `dir\file` در ویندوز.
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// برای ساختن مسیرها به شکل قابل‌حمل باید از `Join`
	// استفاده کرد. این تابع هر تعداد آرگومان می‌گیرد و از
	// آن‌ها یک مسیر سلسله‌مراتبی می‌سازد.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// همیشه به‌جای چسباندن دستی `/` یا `\`، از `Join`
	// استفاده کنید. `Join` علاوه بر قابل‌حمل بودن، مسیرها
	// را با حذف جداکننده‌های اضافی و تغییر پوشه‌های
	// غیرضروری نرمال‌سازی هم می‌کند.
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// با `Dir` و `Base` می‌توان یک مسیر را به پوشه و فایل
	// تقسیم کرد. همچنین `Split` هر دو را در یک فراخوانی
	// برمی‌گرداند.
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// می‌توانیم بررسی کنیم که آیا یک مسیر مطلق است.
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// برخی نام‌های فایل بعد از یک نقطه پسوند دارند. با
	// `Ext` می‌توانیم پسوند را از این نام‌ها جدا کنیم.
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// برای گرفتن نام فایل بدون پسوند، از
	// `strings.TrimSuffix` استفاده کنید.
	fmt.Println(strings.TrimSuffix(filename, ext))

	// `Rel` مسیر نسبی بین یک *مبنا* (base) و یک *مقصد*
	// (target) را پیدا می‌کند. اگر نتوان مقصد را نسبت به
	// مبنا نوشت، خطا برمی‌گرداند.
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}

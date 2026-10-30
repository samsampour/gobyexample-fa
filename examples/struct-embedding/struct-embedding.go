// Go از _جاسازی_ (embedding) ساختارها و اینترفیس‌ها
// پشتیبانی می‌کند تا _ترکیب_ نوع‌ها روان‌تر بیان شود.
// این را با [`//go:embed`](embed-directive) اشتباه نگیرید
// که یک دستور (directive) در Go نسخهٔ 1.16 به بعد است و
// برای جاسازی فایل‌ها و پوشه‌ها در فایل اجرایی برنامه
// به کار می‌رود.

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// یک `container`، ساختار `base` را _جاسازی_ می‌کند.
// جاسازی شبیه یک فیلد بدون نام است.
type container struct {
	base
	str string
}

func main() {

	// وقتی ساختارها را با مقدار لفظی می‌سازیم، باید
	// جاسازی را صریحاً مقداردهی کنیم؛ اینجا نوعِ جاسازی‌شده
	// نقش نام فیلد را دارد.
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// می‌توانیم مستقیماً روی `co` به فیلدهای base دسترسی
	// داشته باشیم، مثلاً `co.num`.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// یا می‌توانیم مسیر کامل را با نام نوع جاسازی‌شده
	// بنویسیم.
	fmt.Println("also num:", co.base.num)

	// چون `container` ساختار `base` را جاسازی کرده، متدهای
	// `base` هم متدهای `container` می‌شوند. اینجا متدی را
	// که از `base` جاسازی شده، مستقیماً روی `co` صدا
	// می‌زنیم.
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// جاسازی ساختارهایی که متد دارند می‌تواند برای دادن
	// پیاده‌سازی اینترفیس به ساختارهای دیگر استفاده شود.
	// اینجا می‌بینیم که `container` حالا اینترفیس
	// `describer` را پیاده‌سازی می‌کند، چون `base` را
	// جاسازی کرده است.
	var d describer = co
	fmt.Println("describer:", d.describe())
}

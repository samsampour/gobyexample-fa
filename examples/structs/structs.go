// _ساختارها_ (structs) در Go مجموعه‌هایی نوع‌دار از
// فیلدها هستند. برای گروه‌بندی داده‌ها کنار هم و ساختن
// رکورد به کار می‌آیند.

package main

import "fmt"

// این نوع ساختار `person` دو فیلد `name` و `age` دارد.
type person struct {
	name string
	age  int
}

// `newPerson` یک ساختار person جدید با نام داده‌شده می‌سازد.
func newPerson(name string) *person {
	// Go زبانی با زباله‌روب (garbage collector) است؛ پس
	// می‌توانید با خیال راحت اشاره‌گری به یک متغیر محلی
	// برگردانید. این متغیر فقط زمانی توسط زباله‌روب پاک
	// می‌شود که هیچ ارجاع فعالی به آن نمانده باشد.
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// این نحو یک ساختار جدید می‌سازد.
	fmt.Println(person{"Bob", 20})

	// هنگام مقداردهی ساختار می‌توانید نام فیلدها را بنویسید.
	fmt.Println(person{name: "Alice", age: 30})

	// فیلدهایی که ذکر نشوند مقدار صفر می‌گیرند.
	fmt.Println(person{name: "Fred"})

	// پیشوند `&` یک اشاره‌گر به ساختار می‌دهد.
	fmt.Println(&person{name: "Ann", age: 40})

	// روش رایج این است که ساختن ساختار جدید را در توابع سازنده (constructor) کپسوله کنیم.
	fmt.Println(newPerson("Jon"))

	// با نقطه به فیلدهای ساختار دسترسی پیدا کنید.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// ساختارها تغییرپذیرند.
	s.age = 51
	fmt.Println(s)

	// با اشاره‌گر به ساختار هم می‌توانید از نقطه استفاده
	// کنید؛ اشاره‌گرها به‌طور خودکار dereference می‌شوند.
	sp := &s
	sp.age = 52
	fmt.Println(sp.age)

	// اگر یک نوع ساختار فقط برای یک مقدار استفاده می‌شود،
	// لازم نیست برایش نام بگذاریم. مقدار می‌تواند نوع
	// ساختارِ بی‌نام داشته باشد. این تکنیک معمولاً در
	// [تست‌های جدول‌محور](testing-and-benchmarking) به کار می‌رود.
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}

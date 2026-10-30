// Go چند تابع مفید برای کار با *پوشه‌ها* (directories)
// در سیستم فایل دارد.

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// یک زیرپوشهٔ جدید در پوشهٔ کاری فعلی می‌سازیم.
	err := os.Mkdir("subdir", 0755)
	check(err)

	// هنگام ساختن پوشه‌های موقت، بهتر است حذفشان را با
	// `defer` زمان‌بندی کنیم. `os.RemoveAll` کل درخت یک
	// پوشه را حذف می‌کند (مشابه `rm -rf`).
	defer os.RemoveAll("subdir")

	// تابع کمکی برای ساختن یک فایل خالی جدید.
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// با `MkdirAll` می‌توانیم سلسله‌مراتبی از پوشه‌ها، از
	// جمله پوشه‌های والد، بسازیم. این مشابه دستور
	// `mkdir -p` در خط فرمان است.
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// `ReadDir` محتوای پوشه را فهرست می‌کند و اسلایسی از
	// اشیای `os.DirEntry` برمی‌گرداند.
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// `Chdir` اجازه می‌دهد پوشهٔ کاری فعلی را عوض کنیم،
	// مشابه `cd`.
	err = os.Chdir("subdir/parent/child")
	check(err)

	// حالا وقتی پوشهٔ *فعلی* را فهرست کنیم، محتوای
	// `subdir/parent/child` را می‌بینیم.
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// با `cd` به جایی که شروع کرده بودیم برمی‌گردیم.
	err = os.Chdir("../../..")
	check(err)

	// همچنین می‌توانیم یک پوشه را به‌صورت *بازگشتی*، همراه با
	// همهٔ زیرپوشه‌هایش، پیمایش کنیم. `WalkDir` یک تابع
	// callback می‌گیرد تا برای هر فایل یا پوشه‌ای که بازدید
	// می‌شود صدا زده شود.
	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit)
	check(err)
}

// `visit` برای هر فایل یا پوشه‌ای که `filepath.WalkDir`
// به‌صورت بازگشتی پیدا می‌کند صدا زده می‌شود.
func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}

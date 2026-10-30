// `//go:embed` یک
// [دستور کامپایلر](https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives)
// است که به برنامه‌ها اجازه می‌دهد فایل‌ها و پوشه‌های دلخواه را هنگام
// بیلد در فایل اجرایی Go قرار دهند. دربارهٔ دستور embed
// [اینجا](https://pkg.go.dev/embed) بیشتر بخوانید.
package main

// بستهٔ `embed` را وارد کنید؛ اگر از هیچ شناسهٔ exported این بسته
// استفاده نمی‌کنید، می‌توانید با `_ "embed"` یک import خالی انجام دهید.
import (
	"embed"
)

// دستورهای `embed` مسیرهایی نسبت به پوشهٔ حاوی فایل کد Go
// می‌پذیرند. این دستور محتوای فایل را در متغیر `string` که
// بلافاصله بعد از آن آمده قرار می‌دهد.
//
//go:embed folder/single_file.txt
var fileString string

// یا محتوای فایل را در یک `[]byte` قرار دهید.
//
//go:embed folder/single_file.txt
var fileByte []byte

// با wildcardها می‌توانیم چند فایل یا حتی پوشه‌ها را هم جاسازی کنیم.
// این کار از متغیری از [نوع embed.FS](https://pkg.go.dev/embed#FS)
// استفاده می‌کند که یک سیستم فایل مجازی ساده را پیاده‌سازی می‌کند.
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	// محتوای `single_file.txt` را چاپ می‌کنیم.
	print(fileString)
	print(string(fileByte))

	// چند فایل را از پوشهٔ جاسازی‌شده می‌خوانیم.
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}

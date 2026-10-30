// Go پشتیبانی داخلی از
// [کدگذاری/رمزگشایی base64](https://en.wikipedia.org/wiki/Base64)
// دارد.

package main

// این نحو بستهٔ `encoding/base64` را به‌جای نام پیش‌فرض
// `base64` با نام `b64` وارد (import) می‌کند. این کار در
// ادامه کمی در جا صرفه‌جویی می‌کند.
import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	// این `string`ی است که آن را کدگذاری/رمزگشایی می‌کنیم.
	data := "abc123!?$*&()'-=@~"

	// Go هم base64 استاندارد و هم base64 سازگار با URL را
	// پشتیبانی می‌کند. این روش کدگذاری با کدگذار استاندارد
	// است. کدگذار یک `[]byte` می‌خواهد، پس `string` خودمان
	// را به آن نوع تبدیل می‌کنیم.
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// رمزگشایی ممکن است خطا برگرداند؛ اگر از درست بودن
	// ورودی مطمئن نیستید، می‌توانید آن را بررسی کنید.
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// این کد با قالب base64 سازگار با URL کدگذاری/رمزگشایی
	// می‌کند.
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}

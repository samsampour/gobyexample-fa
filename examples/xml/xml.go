// Go با بستهٔ `encoding/xml` پشتیبانی داخلی از XML و
// قالب‌های شبه‌XML دارد.

package main

import (
	"encoding/xml"
	"fmt"
)

// Plant به XML نگاشت می‌شود. مشابه مثال‌های JSON، تگ‌های
// فیلدها دستورهایی برای کدگذار و رمزگشا دارند. اینجا از
// چند قابلیت ویژهٔ بستهٔ XML استفاده می‌کنیم: نام فیلد
// `XMLName` نام عنصر XML نمایندهٔ این ساختار را تعیین
// می‌کند؛ `id,attr` یعنی فیلد `Id` یک _ویژگی_ (attribute)
// در XML است، نه یک عنصر تودرتو.
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// XML نمایندهٔ گیاه‌مان را تولید می‌کنیم؛ با
	// `MarshalIndent` تا خروجی خواناتری برای انسان تولید
	// شود.
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// برای افزودن یک سرآیند (header) عمومی XML به خروجی،
	// آن را صریحاً اضافه کنید.
	fmt.Println(xml.Header + string(out))

	// از `Unmarshal` برای تجزیهٔ جریانی از بایت‌های XML به
	// یک ساختار داده استفاده کنید. اگر XML نادرست باشد یا
	// نتوان آن را به Plant نگاشت کرد، یک خطای توصیفی
	// برگردانده می‌شود.
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// تگ فیلد `parent>child>plant` به کدگذار می‌گوید که
	// همهٔ `plant`ها را زیر `<parent><child>...` قرار دهد.
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}

// Go پشتیبانی داخلی از کدگذاری و رمزگشایی JSON دارد،
// از جمله تبدیل از/به انواع داده‌ی داخلی و سفارشی.

package main

import (
	"bytes"
	"encoding/json/v2"
	"fmt"
	"strings"
)

// از این دو ساختار برای نشان دادن کدگذاری و رمزگشایی
// انواع سفارشی در ادامه استفاده می‌کنیم.
type response1 struct {
	Page   int
	Fruits []string
}

// فقط فیلدهای exported در JSON کدگذاری/رمزگشایی می‌شوند.
// فیلدها برای exported بودن باید با حرف بزرگ شروع شوند.
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// اول کدگذاری انواع داده‌ی پایه به رشته‌های JSON را
	// می‌بینیم. این چند مثال برای مقادیر اتمی (ساده) است.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// و این چند مثال برای اسلایس‌ها و مپ‌هاست که همان‌طور
	// که انتظار دارید به آرایه‌ها و آبجکت‌های JSON تبدیل
	// می‌شوند.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// بستهٔ JSON می‌تواند انواع داده‌ی سفارشی شما را خودکار
	// کدگذاری کند. فقط فیلدهای exported را در خروجی
	// می‌آورد و به‌طور پیش‌فرض از نام همان فیلدها به‌عنوان
	// کلیدهای JSON استفاده می‌کند.
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// با تگ‌ها (tags) روی تعریف فیلدهای ساختار می‌توانید
	// نام کلیدهای JSON را سفارشی کنید. برای دیدن نمونه‌ای
	// از این تگ‌ها، تعریف `response2` را در بالا ببینید.
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// حالا رمزگشایی داده‌ی JSON به مقادیر Go را ببینیم. این
	// مثالی برای یک ساختار داده‌ی عمومی است.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// باید متغیری فراهم کنیم که بستهٔ JSON داده‌ی
	// رمزگشایی‌شده را در آن بگذارد. این
	// `map[string]any` یک مپ از رشته‌ها به انواع داده‌ی
	// دلخواه نگه می‌دارد.
	var dat map[string]any

	// این رمزگشایی واقعی است، همراه با بررسی خطاهای
	// احتمالی.
	// برای کوتاه‌تر شدن، در این مثال‌ها خطاها را نادیده
	// می‌گیریم؛ در کد واقعی همیشه باید خطاها را بررسی کنید
	// و متناسب با آن‌ها عمل کنید.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// برای استفاده از مقادیر مپ رمزگشایی‌شده، باید آن‌ها را
	// به نوع مناسبشان تبدیل کنیم. برای مثال اینجا مقدار
	// `num` را به نوع مورد انتظار `float64` تبدیل می‌کنیم.
	num := dat["num"].(float64)
	fmt.Println(num)

	// دسترسی به داده‌های تودرتو نیاز به یک سری تبدیل دارد.
	strs := dat["strs"].([]any)
	str1 := strs[0].(string)
	fmt.Println(str1)

	// می‌توانیم JSON را به انواع داده‌ی سفارشی هم رمزگشایی
	// کنیم. مزیت این کار، ایمنی نوع بیشتر در برنامه و
	// حذف نیاز به type assertion هنگام دسترسی به داده‌ی
	// رمزگشایی‌شده است.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	_ = json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// در مثال‌های بالا همیشه از بایت‌ها و رشته‌ها به‌عنوان
	// واسطه بین داده و نمایش JSON در خروجی استاندارد
	// استفاده کردیم. می‌توانیم کدگذاری JSON را مستقیماً به
	// `io.Writer`هایی مثل `os.Stdout` یا حتی بدنهٔ پاسخ‌های
	// HTTP هم بفرستیم (stream).
	d := map[string]int{"apple": 5, "lettuce": 7}
	var buf bytes.Buffer
	_ = json.MarshalWrite(&buf, d)
	fmt.Println(buf.String())

	// خواندن جریانی از `io.Reader`هایی مثل `os.Stdin` یا
	// بدنهٔ درخواست‌های HTTP با `json.UnmarshalRead` انجام
	// می‌شود.
	res1 := response2{}
	_ = json.UnmarshalRead(strings.NewReader(str), &res1)
	fmt.Println(res1)
}

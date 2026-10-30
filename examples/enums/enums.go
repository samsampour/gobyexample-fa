// _نوع‌های شمارشی_ (enums) حالت خاصی از
// [نوع‌های جمعی](https://en.wikipedia.org/wiki/Algebraic_data_type)
// (sum types) هستند. enum نوعی است که تعداد ثابتی مقدار
// ممکن دارد و هر مقدار نام مشخصی دارد. Go نوع enum را
// به‌عنوان یک قابلیت جداگانهٔ زبان ندارد، اما پیاده‌سازی
// enum با الگوهای رایج زبان ساده است.

package main

import "fmt"

// نوع enum ما، `ServerState`، نوع پایهٔ `int` دارد.
type ServerState int

// مقادیر ممکن `ServerState` به‌صورت ثابت تعریف می‌شوند.
// کلیدواژهٔ ویژهٔ [iota](https://go.dev/ref/spec#Iota)
// مقادیر ثابت متوالی را خودکار تولید می‌کند؛ در این
// مثال 0، 1، 2 و الی آخر.
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// با پیاده‌سازی اینترفیس [fmt.Stringer](https://pkg.go.dev/fmt#Stringer)،
// مقادیر `ServerState` قابل چاپ یا تبدیل به رشته
// می‌شوند.
//
// اگر تعداد مقادیر ممکن زیاد باشد، این کار می‌تواند
// دست‌وپاگیر شود. در چنین مواردی می‌توان از
// [ابزار stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
// همراه با `go:generate` برای خودکارسازی این فرایند
// استفاده کرد. برای توضیح مفصل‌تر [این مقاله](https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate)
// را ببینید.
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	// اگر مقداری از نوع `int` داشته باشیم، نمی‌توانیم آن را به
	// `transition` بدهیم؛ کامپایلر از ناسازگاری نوع خطا
	// می‌گیرد. این کار تا حدی ایمنی نوع در زمان کامپایل
	// را برای enumها فراهم می‌کند.

	ns2 := transition(ns)
	fmt.Println(ns2)
}

// transition یک تغییر وضعیت را برای یک سرور شبیه‌سازی
// می‌کند؛ وضعیت فعلی را می‌گیرد و وضعیت جدید را
// برمی‌گرداند.
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		// فرض کنید اینجا شرط‌هایی را بررسی می‌کنیم تا
		// وضعیت بعدی را تعیین کنیم...
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}

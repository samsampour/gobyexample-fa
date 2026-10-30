// گاهی برنامه‌های Go ما باید پروسه‌های دیگری را ایجاد
// (spawn) کنند.

package main

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// با یک دستور ساده شروع می‌کنیم که آرگومان یا ورودی
	// نمی‌گیرد و فقط چیزی در stdout چاپ می‌کند. تابع کمکی
	// `exec.Command` شیئی می‌سازد که نمایندهٔ این پروسهٔ
	// خارجی است.
	dateCmd := exec.Command("date")

	// متد `Output` دستور را اجرا می‌کند، منتظر پایانش
	// می‌ماند و خروجی استانداردش را جمع می‌کند.
	// اگر خطایی نبود، `dateOut` بایت‌های اطلاعات تاریخ را
	// نگه می‌دارد.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// `Output` و دیگر متدهای `Command` اگر در اجرای دستور
	// مشکلی پیش بیاید (مثلاً مسیر اشتباه) `*exec.Error`
	// برمی‌گردانند، و اگر دستور اجرا شود اما با کد بازگشتی
	// غیرصفر خارج شود، `*exec.ExitError` برمی‌گردانند.
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		if e, ok := errors.AsType[*exec.Error](err); ok {
			fmt.Println("failed executing:", e)
		} else if e, ok := errors.AsType[*exec.ExitError](err); ok {
			exitCode := e.ExitCode()
			fmt.Println("command exit rc =", exitCode)
		} else {
			panic(err)
		}
	}

	// در ادامه حالت کمی پیچیده‌تری را می‌بینیم که در آن
	// داده را از طریق `stdin` پروسهٔ خارجی به آن می‌دهیم و
	// نتایج را از `stdout` آن جمع می‌کنیم.
	grepCmd := exec.Command("grep", "hello")

	// اینجا صریحاً لوله‌های (pipes) ورودی/خروجی را
	// می‌گیریم، پروسه را شروع می‌کنیم، مقداری ورودی در آن
	// می‌نویسیم، خروجی حاصل را می‌خوانیم و در نهایت منتظر
	// خروج پروسه می‌مانیم.
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// در مثال بالا بررسی خطاها را حذف کردیم، اما می‌توانید
	// برای همهٔ آن‌ها از الگوی معمول `if err != nil` استفاده
	// کنید. همچنین فقط نتایج `StdoutPipe` را جمع کردیم،
	// اما می‌توانید `StderrPipe` را هم دقیقاً به همین روش
	// جمع کنید.
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// توجه کنید که هنگام ایجاد دستورها باید دستور و آرایهٔ
	// آرگومان‌ها را صریحاً و جداگانه بدهیم؛ نمی‌توانیم فقط
	// یک رشتهٔ خط فرمان بدهیم. اگر می‌خواهید یک دستور کامل
	// را با یک رشته اجرا کنید، می‌توانید از گزینهٔ `-c` در
	// `bash` استفاده کنید:
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

// در مثال قبل از قفل‌گذاری صریح با [میوتکس‌ها](mutexes)
// برای همگام‌سازی دسترسی به وضعیت مشترک بین چند گوروتین
// استفاده کردیم. گزینهٔ دیگر این است که از قابلیت‌های
// همگام‌سازی داخلی گوروتین‌ها و کانال‌ها برای رسیدن به
// همان نتیجه استفاده کنیم. این رویکرد مبتنی بر کانال با
// ایدهٔ Go هماهنگ است: به اشتراک گذاشتن حافظه از طریق
// ارتباط، و اینکه هر تکه داده دقیقاً متعلق به یک
// گوروتین باشد.

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// در این مثال وضعیت ما متعلق به یک گوروتین واحد است. این
// تضمین می‌کند که داده‌ها هرگز با دسترسی هم‌زمان خراب
// نشوند. برای خواندن یا نوشتن آن وضعیت، گوروتین‌های دیگر
// به گوروتینِ مالک پیام می‌فرستند و پاسخ متناظر را
// دریافت می‌کنند. این `struct`های `readOp` و `writeOp`
// آن درخواست‌ها و راهی برای پاسخ دادن گوروتین مالک را
// در بر می‌گیرند.
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// مثل قبل، تعداد عملیاتی را که انجام می‌دهیم می‌شماریم.
	var readOps uint64
	var writeOps uint64

	// گوروتین‌های دیگر از کانال‌های `reads` و `writes`
	// به‌ترتیب برای ارسال درخواست‌های خواندن و نوشتن
	// استفاده می‌کنند.
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// این گوروتینی است که مالک `state` است؛ `state` مثل
	// مثال قبل یک مپ است، اما حالا مختص همین گوروتینِ
	// حالت‌دار است. این گوروتین پشت سر هم روی کانال‌های
	// `reads` و `writes`، select می‌کند و به درخواست‌ها به
	// محض رسیدن پاسخ می‌دهد. پاسخ دادن به این صورت است که
	// اول عملیات درخواستی انجام می‌شود و سپس مقداری روی
	// کانال پاسخ `resp` فرستاده می‌شود تا موفقیت را نشان دهد
	// (و در مورد `reads`، مقدار خواسته‌شده را).
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// این بخش ۱۰۰ گوروتین راه می‌اندازد که از طریق کانال
	// `reads` درخواست خواندن به گوروتینِ مالک وضعیت
	// می‌فرستند. هر خواندن شامل ساختن یک `readOp`، فرستادن
	// آن روی کانال `reads` و سپس دریافت نتیجه از کانال
	// `resp` داده‌شده است.
	for range 100 {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// ۱۰ نوشتن را هم با روشی مشابه راه می‌اندازیم.
	for range 10 {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// اجازه می‌دهیم گوروتین‌ها یک ثانیه کار کنند.
	time.Sleep(time.Second)

	// در نهایت، تعداد عملیات را می‌خوانیم و گزارش می‌دهیم.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}

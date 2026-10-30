// در این مثال می‌بینیم که چطور با گوروتین‌ها و کانال‌ها
// یک _استخر کارگر_ (worker pool) پیاده‌سازی کنیم.

package main

import (
	"fmt"
	"time"
)

// این worker (کارگر) است که چند نمونهٔ هم‌روند از آن را
// اجرا می‌کنیم. این کارگرها کارها را از کانال `jobs`
// دریافت می‌کنند و نتایج متناظر را روی `results`
// می‌فرستند. برای شبیه‌سازی یک کار سنگین، به ازای هر کار
// یک ثانیه صبر می‌کنیم.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// برای استفاده از استخر کارگرها باید به آن‌ها کار
	// بفرستیم و نتایجشان را جمع کنیم. برای این کار ۲ کانال
	// می‌سازیم.
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// این بخش ۳ کارگر را راه می‌اندازد که در ابتدا مسدودند،
	// چون هنوز کاری وجود ندارد.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// اینجا ۵ کار (`jobs`) می‌فرستیم و سپس آن کانال را با
	// `close` می‌بندیم تا نشان دهیم کار دیگری نداریم.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// در نهایت همهٔ نتایج کار را جمع می‌کنیم. این کار
	// همچنین تضمین می‌کند که گوروتین‌های کارگر تمام شده‌اند.
	// روش دیگر برای منتظر ماندن چند گوروتین، استفاده از
	// [WaitGroup](waitgroups) است.
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

// _دستورهای switch_ شرط‌ها را در چندین شاخه بیان
// می‌کنند.

package main

import (
	"fmt"
	"time"
)

func main() {

	// یک `switch` ساده.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// می‌توانید چند عبارت را در یک `case` با کاما از هم
	// جدا کنید. در این مثال از حالت اختیاری `default` هم
	// استفاده کرده‌ایم.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// `switch` بدون عبارت، روش دیگری برای بیان منطق
	// if/else است. اینجا همچنین نشان می‌دهیم که عبارت‌های
	// `case` لازم نیست ثابت باشند.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// `switch` نوعی (type switch) به‌جای مقدارها، نوع‌ها
	// را مقایسه می‌کند. با آن می‌توانید نوع یک مقدار
	// اینترفیسی را کشف کنید. در این مثال، متغیر `t` در
	// هر شاخه نوعِ مربوط به همان شاخه را خواهد داشت.
	whatAmI := func(i any) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

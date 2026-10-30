// شرط‌گذاری با `if` و `else` در Go ساده و سرراست
// است.

package main

import "fmt"

func main() {

	// یک مثال ساده.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// می‌توانید دستور `if` را بدون else هم بنویسید.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// عملگرهای منطقی مثل `&&` و `||` اغلب در شرط‌ها
	// به کار می‌آیند.
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// پیش از شرط می‌توان یک دستور نوشت؛ هر متغیری که در
	// این دستور تعریف شود، در شاخهٔ فعلی و همهٔ شاخه‌های
	// بعدی در دسترس است.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// توجه کنید که در Go نیازی به پرانتز دور شرط‌ها نیست،
// اما آکولادها الزامی هستند.

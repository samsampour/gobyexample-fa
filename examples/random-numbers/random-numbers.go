// بستهٔ `math/rand/v2` در Go امکان تولید
// [اعداد شبه‌تصادفی](https://en.wikipedia.org/wiki/Pseudorandom_number_generator)
// را فراهم می‌کند.

package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	// برای مثال، `rand.IntN` یک `int` تصادفی n برمی‌گرداند
	// که `0 <= n < 100`.
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	// `rand.Float64` یک `float64` به نام `f` برمی‌گرداند
	// که `0.0 <= f < 1.0`.
	fmt.Println(rand.Float64())

	// از این می‌توان برای تولید اعداد اعشاری تصادفی در
	// بازه‌های دیگر استفاده کرد، مثلاً `5.0 <= f' < 10.0`.
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// اگر یک بذر (seed) مشخص می‌خواهید، یک `rand.Source`
	// جدید بسازید و آن را به سازندهٔ `New` بدهید. `NewPCG`
	// یک منبع
	// [PCG](https://en.wikipedia.org/wiki/Permuted_congruential_generator)
	// جدید می‌سازد که به بذری از دو عدد `uint64` نیاز دارد.
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}

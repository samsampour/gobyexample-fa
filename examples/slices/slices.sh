# توجه کنید که با اینکه نوع اسلایس‌ها با آرایه‌ها فرق
# دارد، `fmt.Println` آن‌ها را شبیه هم نمایش می‌دهد.
$ go run slices.go
uninit: [] true true
emp: [  ] len: 3 cap: 3
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
t == t2
2d:  [[0] [1 2] [2 3 4]]

# برای جزئیات بیشتر دربارهٔ طراحی و پیاده‌سازی اسلایس‌ها
# در Go، این [مقالهٔ عالی](https://go.dev/blog/slices-intro)
# از تیم Go را بخوانید.

# حالا که آرایه‌ها و اسلایس‌ها را دیدیم، سراغ ساختار
# داده‌ی داخلی کلیدی دیگر Go می‌رویم: مپ‌ها.

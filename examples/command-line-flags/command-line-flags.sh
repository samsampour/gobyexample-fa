# برای آزمایش برنامهٔ فلگ‌های خط فرمان، بهتر است اول آن
# را کامپایل کنید و سپس فایل اجرایی حاصل را مستقیماً
# اجرا کنید.
$ go build command-line-flags.go

# برنامهٔ ساخته‌شده را امتحان کنید؛ اول برای همهٔ فلگ‌ها
# مقدار بدهید.
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# توجه کنید که اگر فلگ‌ها را ننویسید، خودکار مقدار
# پیش‌فرضشان را می‌گیرند.
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# آرگومان‌های موقعیتی انتهایی را می‌توان بعد از فلگ‌ها
# نوشت.
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# توجه کنید که بستهٔ `flag` نیاز دارد همهٔ فلگ‌ها پیش از
# آرگومان‌های موقعیتی بیایند (وگرنه فلگ‌ها هم به‌عنوان
# آرگومان موقعیتی تفسیر می‌شوند).
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# از فلگ‌های `-h` یا `--help` استفاده کنید تا متن
# راهنمای خودکار برنامهٔ خط فرمان را ببینید.
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# اگر فلگی بدهید که در بستهٔ `flag` تعریف نشده، برنامه
# یک پیام خطا چاپ می‌کند و دوباره متن راهنما را نشان
# می‌دهد.
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...

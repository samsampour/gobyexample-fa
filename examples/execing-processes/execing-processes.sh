# وقتی برنامه را اجرا می‌کنیم، `ls` جایگزین آن می‌شود.
$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

# توجه کنید که Go تابع کلاسیک `fork` یونیکس را ندارد.
# البته معمولاً این مشکلی نیست، چون راه‌اندازی
# گوروتین‌ها، ایجاد پروسه‌ها و exec کردن پروسه‌ها بیشتر
# کاربردهای `fork` را پوشش می‌دهند.

# کد نوشتن فایل را اجرا کنید.
$ go run writing-files.go 
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

# سپس محتوای فایل‌های نوشته‌شده را بررسی کنید.
$ cat /tmp/dat1
hello
go
$ cat /tmp/dat2
some
writes
buffered

# در ادامه می‌بینیم که چطور برخی از ایده‌های ورودی/خروجی
# فایل را که الان دیدیم، روی جریان‌های `stdin` و `stdout`
# به کار ببریم.

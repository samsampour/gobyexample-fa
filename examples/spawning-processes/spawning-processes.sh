# برنامه‌های ایجادشده همان خروجی‌ای را برمی‌گردانند که
# اگر مستقیماً از خط فرمان اجرایشان می‌کردیم می‌دیدیم.
$ go run spawning-processes.go 
> date
Thu 05 May 2022 10:10:12 PM PDT

# دستور date فلگ `-x` ندارد، پس با یک پیام خطا و کد
# بازگشتی غیرصفر خارج می‌شود.
command exit rc = 1
> grep hello
hello grep

> ls -a -l -h
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 spawning-processes.go

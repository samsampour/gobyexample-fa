# اجرای برنامه نشان می‌دهد که مقدار `FOO` را که در
# برنامه تنظیم کردیم می‌گیریم، اما `BAR` خالی است.
$ go run environment-variables.go
FOO: 1
BAR: 

# فهرست کلیدهای محیط به سیستم خاص شما بستگی دارد.
TERM_PROGRAM
PATH
SHELL
...
FOO

# اگر اول `BAR` را در محیط تنظیم کنیم، برنامهٔ در حال
# اجرا آن مقدار را می‌گیرد.
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...

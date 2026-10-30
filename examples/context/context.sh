# سرور را در پس‌زمینه اجرا کنید.
$ go run context.go &

# یک درخواست کلاینت به `/hello` را شبیه‌سازی کنید و کمی
# بعد از شروع، Ctrl+C را بزنید تا سیگنال لغو فرستاده
# شود.
$ curl localhost:8090/hello
server: hello handler started
^C
server: context canceled
server: hello handler ended

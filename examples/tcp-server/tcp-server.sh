# سرور TCP را در پس‌زمینه اجرا کنید.
$ go run tcp-server.go &

# با netcat داده بفرستید و پاسخ را دریافت کنید.
$ echo "Hello from netcat" | nc localhost 8090
ACK: HELLO FROM NETCAT

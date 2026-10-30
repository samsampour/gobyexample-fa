# برای اجرای مثال از این دستورها استفاده کنید.
# (توجه: به دلیل محدودیت go playground، این مثال فقط
# روی سیستم خودتان قابل اجراست.)
$ mkdir -p folder
$ echo "hello go" > folder/single_file.txt
$ echo "123" > folder/file1.hash
$ echo "456" > folder/file2.hash

$ go run embed-directive.go
hello go
hello go
123
456

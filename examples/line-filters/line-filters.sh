# برای امتحان فیلتر خطی، اول فایلی با چند خط با حروف
# کوچک بسازید.
$ echo 'hello'   > /tmp/lines
$ echo 'filter' >> /tmp/lines

# سپس با فیلتر خطی، خطوط را با حروف بزرگ بگیرید.
$ cat /tmp/lines | go run line-filters.go
HELLO
FILTER

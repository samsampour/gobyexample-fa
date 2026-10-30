$ go build command-line-subcommands.go 

# اول زیرفرمان foo را اجرا کنید.
$ ./command-line-subcommands foo -enable -name=joe a1 a2
subcommand 'foo'
  enable: true
  name: joe
  tail: [a1 a2]

# حالا bar را امتحان کنید.
$ ./command-line-subcommands bar -level 8 a1
subcommand 'bar'
  level: 8
  tail: [a1]

# اما bar فلگ‌های foo را نمی‌پذیرد.
$ ./command-line-subcommands bar -enable a1
flag provided but not defined: -enable
Usage of bar:
  -level int
    	level

# در ادامه متغیرهای محیطی را می‌بینیم؛ روش رایج دیگری
# برای پارامتری کردن برنامه‌ها.

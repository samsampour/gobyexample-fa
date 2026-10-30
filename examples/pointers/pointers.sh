# `zeroval` مقدار `i` را در `main` تغییر نمی‌دهد، اما
# `zeroptr` تغییر می‌دهد، چون به آدرس حافظهٔ آن متغیر
# ارجاع دارد.
$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100
value at *p: 42
value at *p: 0

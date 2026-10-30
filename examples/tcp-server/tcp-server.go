// بستهٔ `net` ابزارهای لازم برای ساختن آسان سرورهای سوکت
// TCP را فراهم می‌کند.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	// `net.Listen` سرور را روی شبکهٔ داده‌شده (TCP) و
	// آدرس داده‌شده (پورت 8090 روی همهٔ رابط‌ها) راه
	// می‌اندازد.
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal("Error listening:", err)
	}

	// listener را می‌بندیم تا هنگام خروج برنامه پورت آزاد
	// شود.
	defer listener.Close()

	// به‌طور نامحدود حلقه می‌زنیم تا اتصال‌های جدید کلاینت‌ها را بپذیریم.
	for {
		// منتظر یک اتصال می‌مانیم.
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting conn:", err)
			continue
		}

		// اینجا از یک گوروتین برای رسیدگی به اتصال استفاده
		// می‌کنیم تا حلقهٔ اصلی بتواند به پذیرفتن اتصال‌های
		// بیشتر ادامه دهد.
		go handleConnection(conn)
	}
}

// `handleConnection` به یک اتصال کلاینت رسیدگی می‌کند؛
// یک خط متن از کلاینت می‌خواند و پاسخی برمی‌گرداند.
func handleConnection(conn net.Conn) {
	// بستن اتصال، وقتی کارمان با کلاینت تمام شد، منابع را
	// آزاد می‌کند.
	defer conn.Close()

	// با `bufio.NewReader` یک خط داده از کلاینت می‌خوانیم
	// (که با کاراکتر خط جدید تمام می‌شود).
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}

	// یک پاسخ می‌سازیم و به کلاینت برمی‌گردانیم تا ارتباط
	// دوطرفه را نشان دهیم.
	ackMsg := strings.ToUpper(strings.TrimSpace(message))
	response := fmt.Sprintf("ACK: %s\n", ackMsg)
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Printf("Server write error: %v", err)
	}
}

package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	target := "google.com"
	if len(os.Args) > 1 {
		target = os.Args[1]
	}

	addr, err := net.LookupHost(target)
	if err != nil {
		fmt.Printf("❌ Cannot resolve %s\n", target)
		return
	}

	fmt.Printf("🏓 Pinging %s (%s)\n", target, addr[0])
	fmt.Println("==================")

	for i := 0; i < 4; i++ {
		start := time.Now()
		conn, err := net.DialTimeout("ip4:icmp", addr[0]+":0", 2*time.Second)
		elapsed := time.Since(start)

		if err != nil {
			fmt.Printf("  ❌ Request timeout\n")
		} else {
			conn.Close()
			fmt.Printf("  ✅ %dms\n", elapsed.Milliseconds())
		}
		time.Sleep(1 * time.Second)
	}
}
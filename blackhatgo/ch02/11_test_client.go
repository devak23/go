// test_client.go - Simulates ACME making requests through the proxy

package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	fmt.Println("🏢 ACME Employee (Joe) testing proxy access...")
	fmt.Println("")

	// Test 1: Direct access to "cat cam" (this would be blocked by ACME firewall)
	fmt.Println("❌ Attempting direct access to joescatcam.website (localhost:8080)")
	fmt.Println("   (In real scenario, this would be blocked by ACME firewall)")

	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Printf("   Direct access failed: %v\n", err)
	} else {
		resp.Body.Close()
		fmt.Println("   Direct access worked (but would be blocked in real scenario)")
	}
	fmt.Println("")

	// Test 2: Access through proxy (this works!)
	fmt.Println("✅ Attempting access through joesproxy.com (localhost:8081)")
	fmt.Println("   (This bypasses ACME's firewall)")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	for i := 1; i <= 3; i++ {
		fmt.Printf("   Request #%d: ", i)

		resp, err := client.Get("http://localhost:8081")
		if err != nil {
			fmt.Printf("Failed: %v\n", err)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Printf("Error reading response: %v\n", err)
			continue
		}

		fmt.Printf("Success! Received %d bytes\n", len(body))

		if i == 1 {
			fmt.Println("   📄 Response preview:")
			if len(body) > 200 {
				fmt.Printf("   %s...\n", string(body[:200]))
			} else {
				fmt.Printf("   %s\n", string(body))
			}
		}

		time.Sleep(1 * time.Second)
	}

	fmt.Println("")
	fmt.Println("🎉 Proxy test complete!")
	fmt.Println("Joe successfully accessed his cat cam through the proxy!")
}

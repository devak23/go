// test_proxy.go - Complete local testing setup for port forwarding

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

// This simulates joescatcam.website - the "blocked" destination
func startCatCamServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `
<!DOCTYPE html>
<html>
<head><title>Joe's Cat Cam (Simulated)</title></head>
<body style="font-family: Arial; text-align: center; background: #f0f8ff;">
	<h1>🐱 Joe's Cat Cam Website 🐱</h1>
	<p>This simulates joescatcam.website</p>
	<p>You're viewing this through the proxy!</p>
	<div style="font-size: 48px;">😸 😺 😻</div>
	<p><em>Streaming live cats in 4K Ultra HD!</em></p>
	<p>Request came from: ` + r.RemoteAddr + `</p>
	<p>Time: ` + time.Now().Format("15:04:05") + `</p>
</body>
</html>`
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, html)
	})

	log.Println("🐱 Cat cam server starting on :8080 (simulates joescatcam.website)")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Cat cam server failed:", err)
	}
}

// This is Joe's proxy server (joesproxy.com)
func startProxyServer() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("Unable to bind proxy to port 8081:", err)
	}

	log.Println("🔄 Proxy server starting on :8081 (simulates joesproxy.com)")
	log.Println("   Forwarding all traffic to localhost:8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Unable to accept connection:", err)
			continue
		}
		go handleProxy(conn)
	}
}

func handleProxy(src net.Conn) {
	defer src.Close() // Close connection from "ACME" to proxy

	log.Printf("📡 Proxy: New connection from %s", src.RemoteAddr())

	// Connect to our "cat cam" (localhost:8080 instead of joescatcam.website:80)
	dst, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Printf("❌ Proxy: Unable to connect to target: %v", err)
		return
	}
	defer dst.Close() // Close proxy to cat cam connection

	log.Printf("✅ Proxy: Connected to target, starting relay")

	// Copy data bidirectionally
	go func() {
		defer dst.Close()
		if _, err := io.Copy(dst, src); err != nil {
			log.Printf("⬆️  Proxy: Error copying to target: %v", err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Printf("⬇️  Proxy: Error copying from target: %v", err)
	}

	log.Printf("🔚 Proxy: Connection from %s closed", src.RemoteAddr())
}

func main() {
	log.Println("🧪 Starting local port forwarding test environment")
	log.Println("")
	log.Println("This simulates:")
	log.Println("  • localhost:8080 = joescatcam.website (the 'blocked' site)")
	log.Println("  • localhost:8081 = joesproxy.com (the proxy)")
	log.Println("")
	log.Println("Test it by visiting: http://localhost:8081")
	log.Println("")

	// Start the cat cam server in a goroutine
	go startCatCamServer()

	// Give the cat cam server a moment to start
	time.Sleep(1000 * time.Millisecond)

	// Start the proxy server (this blocks)
	startProxyServer()
}

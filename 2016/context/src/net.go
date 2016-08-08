package main

import "net"

func main() {
	// START OMIT
	dialer := &net.Dialer{}
	conn, err := dialer.DialContext(ctx, "example.com", "80") // HL
	// END OMIT
}

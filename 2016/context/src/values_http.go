package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
)

// START OMIT
func handle(w http.ResponseWriter, r *http.Request) {
	server := r.Context().Value(http.ServerContextKey).(*http.Server) // HL
	laddr := r.Context().Value(http.LocalAddrContextKey).(net.Addr)   // HL

	fmt.Printf("Server:%+v\n", server)
	fmt.Printf("LocalAddr:%s\n", laddr.String())
}

// END OMIT

var ts = httptest.NewServer(http.HandlerFunc(handle))

func main() {
	http.Get(ts.URL)
}

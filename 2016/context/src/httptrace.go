package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httptrace"
)

func handle(w http.ResponseWriter, r *http.Request) {
}

var ts = httptest.NewServer(http.HandlerFunc(handle))

func main() {
	log.SetFlags(log.Lmicroseconds)

	// START OMIT
	trace := &httptrace.ClientTrace{
		GetConn:              func(hostPort string) { log.Printf("GetConn:%s", hostPort) },
		GotConn:              func(c httptrace.GotConnInfo) { log.Printf("GotConn:%+v", c) },
		PutIdleConn:          func(err error) { log.Println("PutIdleConn:", err) },
		GotFirstResponseByte: func() { log.Println("GotFirstResponseByte") },
		Got100Continue:       func() { log.Println("Got100Continue") },
		DNSStart:             func(d httptrace.DNSStartInfo) { log.Printf("DNSStart:%+v", d) },
		DNSDone:              func(d httptrace.DNSDoneInfo) { log.Printf("DNSDone:%+v", d) },
		ConnectStart:         func(network, addr string) { log.Println("ConnectStart:", network, addr) },
		ConnectDone:          func(network, addr string, err error) { log.Println("ConnectDone:", network, addr, err) },
		WroteRequest:         func(w httptrace.WroteRequestInfo) { log.Printf("WroteRequest:%+v", w) },
	}

	req, _ := http.NewRequest("GET", ts.URL, nil)
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace)) // tracerをセット // HL
	http.DefaultClient.Do(req)
	// END OMIT
}

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

func slowProcess(ctx context.Context) error { // HL
	for i := 0; i < 10; i++ {
		log.Println("doing something...", i)
		select {
		case <-time.After(1 * time.Second):
		case <-ctx.Done(): // HL
			log.Println("slowProcess done.", i)
			return ctx.Err()
		}
	}
	log.Println("something is done")
	return nil
}

// START OMIT
func handle(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second) // r.Context()を派生 // HL
	defer cancel()                                                 // HL

	resultCh := make(chan error, 1)
	go func() { resultCh <- slowProcess(ctx) }() // HL

	err := <-resultCh
	fmt.Fprintln(w, "Result:", err)
}

// END OMIT

func main() {
	var ts = httptest.NewServer(http.HandlerFunc(handle))

	resp, _ := http.Get(ts.URL)
	b, _ := ioutil.ReadAll(resp.Body)
	log.Printf("%s", b)
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

func slowProcess(done <-chan struct{}) error { // HL
	for i := 0; i < 10; i++ {
		log.Println("doing something...", i)
		select {
		case <-time.After(1 * time.Second):
		case <-done: // HL
			log.Println("slowProcess done.", i)
			return nil // HL
		}
	}
	log.Println("something is done")
	return nil
}

// START OMIT
func handle(w http.ResponseWriter, r *http.Request) {
	done := make(chan struct{}, 1) // 終了を通知するchannel // HL
	resultCh := make(chan error, 1)
	go func() { resultCh <- slowProcess(done) }() // HL

	timer := time.NewTimer(3 * time.Second)
	defer timer.Stop()

	select {
	case <-timer.C:
		close(done) // HL
		fmt.Fprintln(w, "slowProcess Timedout.")
	case err := <-resultCh:
		fmt.Fprintln(w, "Result:", err)
	}
}

// END OMIT

func main() {
	var ts = httptest.NewServer(http.HandlerFunc(handle))

	resp, _ := http.Get(ts.URL)
	b, _ := ioutil.ReadAll(resp.Body)
	log.Printf("%s", b)
}

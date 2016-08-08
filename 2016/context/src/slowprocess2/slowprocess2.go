package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

func slowProcess() error {
	for i := 0; i < 10; i++ {
		log.Println("doing something...", i)
		time.Sleep(time.Duration(1 * time.Second))
	}
	log.Println("something is done")
	return nil
}

// START OMIT
func handle(w http.ResponseWriter, r *http.Request) {
	resultCh := make(chan error, 1)
	go func() { resultCh <- slowProcess() }()

	timer := time.NewTimer(3 * time.Second) // HL
	defer timer.Stop()                      // HL

	select {
	case <-timer.C: // HL
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

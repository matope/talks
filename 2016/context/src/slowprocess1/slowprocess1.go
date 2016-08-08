package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

// START OMIT
func slowProcess() error {
	for i := 0; i < 10; i++ {
		log.Println("doing something...", i)
		time.Sleep(time.Duration(1 * time.Second))
	}
	log.Println("something is done")
	return nil
}

func handle(w http.ResponseWriter, r *http.Request) {
	resultCh := make(chan error, 1)           // 結果を受け取るchannel
	go func() { resultCh <- slowProcess() }() // goroutineで起動

	select {
	case err := <-resultCh: // 結果を待ち受け
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

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

var badServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	<-time.After(100 * time.Millisecond)
	w.WriteHeader(400)
}))

func main() {
	// START OMIT
	req, _ := http.NewRequest("PUT", badServer.URL, strings.NewReader("foo"))
	req.Header.Set("Expect", "100-continue") // HL
	resp, _ := http.DefaultClient.Do(req)

	b, _ := ioutil.ReadAll(req.Body)
	fmt.Printf("Status:%s, Body:[%s]", resp.Status, b)
	// END OMIT
}

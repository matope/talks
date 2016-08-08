package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

var ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	<-time.After(10 * time.Second)
}))

func main() {
	// START OMIT
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req, _ := http.NewRequest("GET", ts.URL, nil)
	req = req.WithContext(ctx) // リクエストのContextを置き換え // HL

	_, err := http.DefaultClient.Do(req)
	fmt.Println(err)
	// END OMIT
}

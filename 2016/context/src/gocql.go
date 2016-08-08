package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

func main() {

	session := &gocql.Session{}

	// START OMIT
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	iter := session.Query(`SELECT id, text FROM tweet WHERE user=?`, "me").WithContext(ctx).Iter() // HL
	for iter.Scan(&id, &text) {
		fmt.Println("Tweet:", id, text)
	}
	// END OMIT
}

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// START OMIT
	ctx, cancel := context.WithCancel(context.Background()) // HL
	go func() {
		<-time.After(1 * time.Second)
		cancel()
	}()

	err := slowProcess(ctx)
	fmt.Println(err)
	// END OMIT
}

func slowProcess(ctx context.Context) error { // HL
	fmt.Println("doing something...")
	select {
	case <-ctx.Done(): // HL
		return ctx.Err()
	}
	return nil
}

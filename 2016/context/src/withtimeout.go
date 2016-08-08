package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	callSlowProcess()
}

func callSlowProcess() {
	// START OMIT
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second) // HL
	defer cancel()

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

package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// START OMIT
	b, err := exec.CommandContext(ctx, "sleep", "100").Output() // HL
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)
	// END OMIT
}

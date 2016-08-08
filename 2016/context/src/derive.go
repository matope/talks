package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// START OMIT
	parent, _ := context.WithTimeout(context.Background(), 1*time.Second)

	// parentを派生
	ctx, _ := context.WithCancel(parent)

	// parent.Done()がクローズすると、派生ContextのDone()もクローズする
	<-ctx.Done()
	fmt.Println(ctx.Err())
	// END OMIT
}

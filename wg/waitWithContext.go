package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var submitCount int32
	// run this instead of wg.Add(1)
	atomic.AddInt32(&submitCount, 1)

	// run this instead of wg.Done()
	// atomic.AddInt32(&submitCount, -1)

	timeout := time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	fmt.Printf("Wait for waitgroup (up to %s)\n", timeout)

	waitWithCtx(ctx, &submitCount)

	fmt.Println("Free at last")
}

// waitWithCtx returns when passed counter drops to zero
// or when context is cancelled
func waitWithCtx(ctx context.Context, counter *int32) {
	ticker := time.NewTicker(10 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if atomic.LoadInt32(counter) == 0 {
				return
			}
		}
	}
}

package terminator

import (
	"context"
	"fmt"
	"time"
)

func ExampleShutDown_background_worker() {
	// --- Job processor ---
	process := func(job string) {
		fmt.Println(job)
	}
	// --- Background Worker ---
	queue := make(chan string, 10)
	worker := func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("worker stopped")
				return
			case job := <-queue:
				process(job)
			}
		}
	}
	// --- Launch the background worker inside a context.WithCancel ---
	workerCtx, workerCancel := context.WithCancel(context.Background())
	workerDone := make(chan struct{})
	go func() {
		defer close(workerDone)
		worker(workerCtx)
	}()
	// --- Shutdown Manager ---
	sm := NewShutdownManager(5 * time.Second)
	// --- Register a drain function for the background worker
	sm.Register(
		"worker-drain",
		20*time.Second,
		func(ctx context.Context) error {
			fmt.Println("draining worker")
			workerCancel() // signal the worker to stop
			select {
			case <-workerDone:
				return nil
			case <-ctx.Done():
				return fmt.Errorf("worker did not drain in time: %w", ctx.Err())
			}
		})
	// Submit jobs to the worker queue
	for i := range 5 {
		queue <- fmt.Sprintf("%s %d", "job", i+1)
	}
	fmt.Println("jobs sent")
	// Give jobs a chance then quit gracefully
	time.Sleep(1 * time.Second)
	if err := sm.Shutdown(); err != nil {
		fmt.Println("shutdown completed with errors", "error", err)
	}

	// Output:
	// jobs sent
	// job 1
	// job 2
	// job 3
	// job 4
	// job 5
	// draining worker
	// worker stopped
}

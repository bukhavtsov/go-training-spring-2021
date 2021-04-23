package main

import (
	"context"
	"fmt"
	"time"
)

// Example №1
//func main() {
//	doJob()
//}
//
//func doJob() {
//	for {
//		fmt.Println("Job is performing...")
//		time.Sleep(time.Second)
//	}
//}


//// Example #2
//func main() {
//	ctx := context.Background()
//	ctx, cancel := context.WithCancel(ctx)
//	go func() {
//		fmt.Scanln()
//		cancel()
//	}()
//
//	doJob(ctx)
//}
//
//func doJob(ctx context.Context) {
//	for {
//		time.Sleep(time.Second)
//		select {
//		case <-ctx.Done():
//			fmt.Println("Job has been stopped")
//			return
//		default:
//			fmt.Println("job is processing...")
//		}
//	}
//}




//// Example #3
//func main() {
//	ctx := context.Background()
//	ctx, cancel := context.WithCancel(ctx)
//	go func() {
//		fmt.Scanln()
//		cancel()
//	}()
//
//	go doJob(ctx, 1)
//	go doJob(ctx, 2)
//	go doJob(ctx, 3)
//
//	time.Sleep(1*time.Minute)
//}
//
//func doJob(ctx context.Context, n int) {
//	for {
//		time.Sleep(time.Second * 2)
//		select {
//		case <-ctx.Done():
//			fmt.Println("Job has been stopped", n)
//			return
//		default:
//			fmt.Println("job is processing...", n)
//		}
//	}
//}


// Example #3
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ctx, cancel = context.WithTimeout(ctx, time.Second * 3)
	defer cancel()

	doJob(ctx, 1)
	fmt.Println("timeout!")
}

func doJob(ctx context.Context, n int) {
	for {
		time.Sleep(time.Second * 1)
		select {
		case <-ctx.Done():
			fmt.Println("Job has been stopped", n)
			return
		default:
			fmt.Println("job is processing...", n)
		}
	}
}

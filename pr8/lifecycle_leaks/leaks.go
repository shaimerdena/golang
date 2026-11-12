
package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
	"sync"
)

// leak 1: nobody takes the result
func slowQuery() int {
	time.Sleep(50 * time.Millisecond) // pretend to ask a database
	return 42
}

func fetchLeaky(timeout time.Duration) (int, error) {
	ch := make(chan int) // unbuffered!

	go func() {
		ch <- slowQuery()	// blocks forever if fethLeaky already returned due to timeout
	}()

	select {
	case res := <-ch:
		return res, nil
	case <-time.After(timeout):
		return 0, errors.New("timeout")
	}
}

// fix 1: give the caller a way to stop the slow query
func fetchFixed(ctx context.Context, timeout time.Duration) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	ch := make(chan int)

	go func() {
		result := slowQuery()

		select {
		case ch <- result:
		case <-ctx.Done():
			return
		}
	}()

	select {
	case res := <-ch:
		return res, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

// fix 2: buffer the channel so the goroutine can finish and exit
func fetchFixedBuffered(timeout time.Duration) (int, error) {
	ch := make(chan int, 1) // buffered channel

	go func() {
		ch <- slowQuery() // blocks forever if fethLeaky already returned due to timeout
	}()

	select {
	case res := <-ch:
		return res, nil
	case <-time.After(timeout):
		return 0, errors.New("timeout")
	}
}

// leak 2: a worker nobody can stop
func startTickerLeaky() {
	go func() {
		for {
			time.Sleep(10 * time.Millisecond) //check "something"
		}
	}()
}

// fix: give the caller a way to stop the ticker
func startTickerFixed(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		ticker := time.NewTicker(10 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				//check "something"
			}
		}
	}()
}


// tasks
// 1. Write handle and buildReport as above. Start 1000 handle calls with a 5
// ms timeout, three rounds, and print goroutines and heap after each
// round (call runtime.GC() first).
func buildReport() *Report {
	time.Sleep(50 * time.Millisecond) // pretend to build a report
	return &Report{}
}

type Report struct {
	Message string
}
func handle(timeout time.Duration) error {
	ch := make(chan *Report)
	go func() { ch <- buildReport() }()

	select {
	case report := <-ch:
		_ = report
	case <-time.After(timeout):
		return errors.New("timeout")
	}
	return nil
}
// 2. Fix it in two steps: (a) a buﬀered channel; (b) pass a
// context.WithTimeout to buildReport so it also stops working when nobody
// needs the result.
// a):
func handleBuffered(timeout time.Duration) error {
	ch := make(chan *Report, 1)
	go func() {ch <- buildReport()} ()
	
	select {
	case report := <-ch:
		_ = report
	case <-time.After(timeout):
		return errors.New("timeout")
	}
	return nil;
}
// b):
func buildReportContext(ctx context.Context) (*Report, error) {
	select {
	case <- time.After(50 * time.Millisecond):
		return &Report{Message: "Report generated"}, nil
	case <- ctx.Done():
		return nil, ctx.Err()
	}
}
func handleContext(timeout time.Duration) error{ 
	ctx, cancel := context.WithTimeout(context.Background(), timeout,)
	defer cancel()

	ch := make(chan *Report, 1)
	go func() {
		report, err := buildReportContext(ctx)
		if err != nil {
			return
		}
		ch <- report
	}()
	select {
	case report := <- ch:
		_ = report
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}

}

func main() {
	startTickerLeaky()

	fmt.Println("Number of goroutines after starting leaky ticker:", runtime.NumGoroutine())

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	startTickerFixed(ctx, &wg)

	fmt.Println("Number of goroutines after starting fixed ticker:", runtime.NumGoroutine())

	cancel()
	wg.Wait()
	time.Sleep(10 * time.Millisecond)

	fmt.Println("After cancel:", runtime.NumGoroutine())

	result, err := fetchLeaky(10 * time.Millisecond)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Result:", result)

	time.Sleep(60 * time.Millisecond)
	fmt.Println("Number of goroutines after leaky fetch:", runtime.NumGoroutine())

	ctx, cancel = context.WithCancel(context.Background())
	result, err = fetchFixed(ctx, 10*time.Millisecond)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Result:", result)

	time.Sleep(60 * time.Millisecond)
	fmt.Println("Number of goroutines after fixed fetch:", runtime.NumGoroutine())
	cancel()

	result, err = fetchFixedBuffered(10 * time.Millisecond)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Result:", result)

	time.Sleep(60 * time.Millisecond)
	fmt.Println("Number of goroutines after fixed buffered fetch:", runtime.NumGoroutine())

	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)	// print goroutine profile to stdout

	// In a real server, import "net/http/pprof" and open /debug/pprof/goroutine?
	// debug=1. In tests, the library go.uber.org/goleak fails a test that leaves
	// goroutines behind. Go 1.26 added an experimental goroutine leak profile,
	// and in Go 1.27 it is generally available ( /debug/pprof/goroutineleak ): the
	// garbage collector finds goroutines blocked on channels or locks that
	// nothing can ever wake up.

	// ch := make(chan int)
	// ch <- 1 	// block forever, leak a goroutine


	// TIPS
	// 1. Before writing go , answer: how and when will this goroutine stop?
	// 2. If a goroutine sends a result that may be ignored, give the channel a
	// buffer large enough for all sends.
	// 3. Long-running goroutines take a context.Context (or a done channel) and
	// return on ctx.Done().
	// 4. Check runtime.NumGoroutine() in tests and look at the goroutine profile
	// in production.


	// tasks
	for round := 1; round <= 3; round++ {
		for call:=0; call<1000; call++ {
			handle(1000)
		}
		fmt.Println(runtime.NumGoroutine())
		
	}
	runtime.GC()
	time.Sleep(100 * time.Millisecond)
	fmt.Println(runtime.NumGoroutine())

	for round := 1; round <= 3; round++ {
		for call:=0; call<1000; call++ {
			handleBuffered(1000)
		}
		fmt.Println(runtime.NumGoroutine())
	}
	runtime.GC()
	time.Sleep(100 * time.Millisecond)
	fmt.Println(runtime.NumGoroutine())

	for round := 1; round <= 3; round++ {
		for call:=0; call<1000; call++ {
			handleContext(1000)
		}
		fmt.Println(runtime.NumGoroutine())
	}
	runtime.GC()
	time.Sleep(100 * time.Millisecond)
	fmt.Println(runtime.NumGoroutine())

	// 3. Show that goroutines and heap stay flat after the fix. Why is step (b)
	// better than step (a) alone?
	// A: bc step (b) allows the goroutine to stop working when nobody needs the result, 
	// while step (a) only buffers the channel and does not stop the goroutine from running.
}
package main
import "fmt"
import "time"
import "sync"

func preemptive() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("goroutine 1 bgein:")
		for i:=0; i<10; i++ {
			fmt.Println("goroutine 1: ", i)
			time.Sleep(1*time.Millisecond)
		}
	} ()

	go func() {
		defer wg.Done()
		fmt.Println("goroutine 2 bgein:")
		for i:=0; i<10; i++ {
			fmt.Println("goroutine 2: ", i)
			time.Sleep(1*time.Millisecond)
		}
	} ()

	wg.Wait()
}


func notpreemptive() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("goroutine 1 bgein:")
		for i:=0; i<1000000000; i++ {
			//fmt.Println("goroutine 1: ", i)
			//time.Sleep(1*time.Millisecond)
		}
	} ()

	go func() {
		defer wg.Done()
		fmt.Println("goroutine 2 bgein:")
		for i:=0; i<1000000000; i++ {
			//fmt.Println("goroutine 2: ", i)
			//time.Sleep(1*time.Millisecond)
		}
	} ()

	wg.Wait()
}

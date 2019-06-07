package main
import "sync"
import "fmt"

func testMutex () {
	var count int
	var lock sync.Mutex

	increment := func() {
		defer lock.Unlock()
		lock.Lock()
		//defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		defer lock.Unlock()
		lock.Lock()
		//defer lock.Unlock()
		count --
		fmt.Printf("Decrementing: %d\n", count)
	}

	var arithmetic sync.WaitGroup
	for i:=0; i<5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		} ()
	}

	for i:=0; i<5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		} ()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete")
}

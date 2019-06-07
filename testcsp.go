package main
import "sync"
import "fmt"
import "time"

func testCsp() {
	var wg sync.WaitGroup
	var i int
	var salutation string
	for i, salutation = range []string {"hello", "greetings", "good day"} {
		wg.Add(1)
		fmt.Println(i, " ", &i, " ", salutation, " ", &salutation)
		go func() {
			defer wg.Done()
			fmt.Println(i, " ", &i, " ", salutation, " ", &salutation)
		} ()
	}
	time.Sleep(10*time.Second)
	wg.Wait()

}

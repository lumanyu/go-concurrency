package main

import "fmt"
import "sync"
import "time"
import "math"
import "text/tabwriter"
import "os"

func testrwmutex() {
	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		for i:=5; i>0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(1)
			//time.Sleep(200*time.Millisecond)
			time.Sleep(200*time.Millisecond)
		}
	}

	consumer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		l.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration{
		var wg sync.WaitGroup
		wg.Add(count+1)
		beginTestTime := time.Now()
		go producer(&wg, mutex)

		for i:= count; i>0; i-- {
			go consumer(&wg, rwMutex)
		}
		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutext\tMutex\n")

	for i:=0; i<15; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m),
		)
	}
}

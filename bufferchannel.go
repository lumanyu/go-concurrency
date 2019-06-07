package main

import "fmt"
import "os"
import "bytes"

func bufferchannel () {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done")
		//for i:=0; i<5; i++ {
		for i:=0; i<10; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <-i
		}
	} ()

	for interger := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", interger)
	}
}

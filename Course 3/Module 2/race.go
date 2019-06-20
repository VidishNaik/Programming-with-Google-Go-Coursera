/* threads share global variables and instructions can be interleaved in any order if multiple threads access same global value. we can not the correct instruction sequence.
For this example, thread 2 may start first and print then change to thread 1 and then back to thread 2 and do i++, so we may see thread2 print 0 and thread 1 print 0, we can try
many times and may see may different results*/


package main

import (
	"fmt"
	"time"
)

func main() {
        var i = 0
	var i_limit = 30
        go func() {
                for ; i < i_limit; i++ {
                        fmt.Println("Thread 1")
			time.Sleep(1 * time.Second)
                }
        }()

	go func() {
                for ; i < i_limit; i++ {
                        fmt.Println("Thread 2")
			time.Sleep(1 * time.Second)
                }
        }()

        for ; i < i_limit; i++{
                fmt.Println("Thread 3")
		time.Sleep(1 * time.Second)
        }
}

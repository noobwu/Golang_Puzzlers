package main

import (
	"fmt"
	"sync/atomic"
	"time"

	"golang.org/x/sys/windows"
)

func main() {
	var count uint32
	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}
	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Printf("i:%d,currentThreadId:%d\n", i, windows.GetCurrentThreadId())
			}
			trigger(i, fn)
		}(i)
	}
	trigger(10, func() {})
}

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 229.

// Pipeline2 demonstrates a finite 3-stage pipeline.
package main

import (
	"fmt"
	"time"
)

// !+
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 1; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		// for x := range naturals {
		// 	squares <- x * x
		// }
		now := time.Now()
		for {
			x, ok := <-naturals
			fmt.Println(ok)
			squares <- x * x
			if time.Since(now).Seconds() > 1 {
				break
			}
		}

		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}

//!-

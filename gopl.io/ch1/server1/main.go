// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 19.
//!+

// Server1 is a minimal "echo" server.
package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("aaaaaa")
	fmt.Printf("%+v", err)
}

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1:])
}
// Every executable Go Program should contain a package called main.
// This tells the Go compiler to compile the package into an executable
// program rather than a shared library.
package main

import (
	"fmt"
	"strings"
)

// The entry point of a Go program should be the main function of main package.
// When the executable is run, main() is automatically called.
func main() {
	fmt.Println("Hello World\n")

	input := "There once was a cat named Barry. He was a very good cat. This cat lived in Boston. He loved doing Boston-related activities (that were good for cats). He walked the esplanade. He shopped on Newbury. He ate at Tatte. He sometimes even went to TD Garden. Did you know that cats are not allowed in TD Garden?"
	substr := input
	index := 0

	for strings.Index(substr, "cat") != -1 {
		index += strings.Index(substr, "cat")

		fmt.Println("found cat @", index)

		_, substr, _ = strings.Cut(substr, "cat")
		index += 3
	}
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// ex 1.1
	// var s, sep string
	// for i := 0; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	// fmt.Println(s)

	// ex 1.2
	// for i, v := range os.Args {
	// 	fmt.Printf("index: %d -> value: %s\n", i, v)
	// }

	// ex 1.3

}

func PrintArgsUsingSeparator() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func PrintArgsUsingJoin() {
	fmt.Println(strings.Join(os.Args, " "))
}

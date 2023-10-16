// ./learn -word="hello", -number=23
package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("number", 42, "a number")

	fmt.Println("the word is : ", *wordPtr)
	fmt.Println("the number is: ", *numPtr)

}

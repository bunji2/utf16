package main

import (
	"fmt"
	"os"

	"github.com/bunji2/utf16"
)

func main() {
	os.Exit(run())
}

func run() int {

	s, b, err := utf16.ReadFileUTF16(os.Args[1])
	if err != nil {
		panic(err)
	}

	printHex(b)
	fmt.Printf("[%s]\n", s)
	return 0
}

func printHex(bs []byte) {
	for _, b := range bs {
		fmt.Printf("%02X,", b)
	}
	fmt.Println("")
}

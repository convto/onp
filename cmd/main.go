package main

import (
	"fmt"
	"os"

	"github.com/convto/onp"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("wrong count of args")
		os.Exit(1)
	}
	d := onp.NewDiff([]rune(args[1]), []rune(args[2]))
	fmt.Printf("edit distance: %d\n", d.ONP())
}

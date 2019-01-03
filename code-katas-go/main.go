package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cluckerino/code-katas-go/hello"
)

func main() {
	fmt.Println(hello.World())
	fmt.Println("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

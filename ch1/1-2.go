package main

import "os"
import "fmt"

func main() {
	for i, arg := range os.Args[:] {
		fmt.Printf("%v %s\n", i, arg)
	}
}

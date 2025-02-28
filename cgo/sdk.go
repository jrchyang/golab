package main

import "C"
import "fmt"

//export hello_from_go
func hello_from_go(name *C.char) {
	fmt.Printf("Go says: Hello, %s!\n", C.GoString(name))
}

func main() {}

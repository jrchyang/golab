package basic

import (
	"fmt"
	"os"
	"testing"
)

func Test_Echo1(t *testing.T) {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func Test_Echo2(t *testing.T) {
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
	}
}

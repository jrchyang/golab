package basic

import (
	"fmt"
	"testing"

	"github.com/jrchyang/golab/utils"
)

func Test_Echo1(t *testing.T) {
	var s, sep string
	args := utils.ParseTestProgramArgs()
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func Test_Echo2(t *testing.T) {
	args := utils.ParseTestProgramArgs()
	for _, arg := range args {
		fmt.Println(arg)
	}

}

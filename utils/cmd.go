package utils

import "os"

func ParseTestProgramArgs() []string {
	args := os.Args
	if len(args) == 0 {
		return []string{}
	}

	result := []string{args[0]}
	index := -1
	for i, arg := range args {
		if arg == "--" {
			index = i
			break
		}
	}

	if index != -1 && index+1 < len(args) {
		result = append(result, args[index+1:]...)
	}

	return result
}

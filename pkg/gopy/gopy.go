package gopy

import (
	"strconv"
)

func searchAndDeletePackage(input string) (count int) {
	for _, c := range input {
		_, err := strconv.Atoi(string(c))
		if err == nil {
			count++
		}
	}
	return count
}

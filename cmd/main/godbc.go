package main

import (
	"fmt"

	"github.com/lpabon/godbc"
)

func Divide(a, b int) int {
	godbc.Require(b != 0, "b should not be 0")
	return a / b
}

func GoDbcRun() {
	fmt.Println(Divide(1, 0))
}

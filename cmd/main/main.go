package main

import (
	"fmt"
	"time"
)

func main() {
	// pkgimports.GetHostInfo()
	l, _ := time.LoadLocation("America/New_York")
	fmt.Printf("%v\n", *l)
}

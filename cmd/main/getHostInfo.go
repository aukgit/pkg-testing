package main

import (
	"fmt"

	"github.com/shirou/gopsutil/host"
)

func getHostInfo(){
	fmt.Println(host.PlatformInformation().)
}

package pkgimports

import (
	"fmt"

	"github.com/shirou/gopsutil/host"
)

func GetHostInfo() {
	fmt.Println(host.PlatformInformation())
}

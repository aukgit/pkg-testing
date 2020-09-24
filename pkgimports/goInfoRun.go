package pkgimports

import (
	"github.com/matishsiao/goInfo"
)

func GoInfoRun() {
	gi := goInfo.GetInfo()
	gi.VarDump()
}

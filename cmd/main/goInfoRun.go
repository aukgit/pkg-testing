package main

import (
	"github.com/matishsiao/goInfo"
)

func goInfoRun() {
	gi := goInfo.GetInfo()
	gi.VarDump()
}

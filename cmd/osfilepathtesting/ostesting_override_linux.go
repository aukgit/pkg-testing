package main

var overrideVariable = "linux.alim.override"

func OverrideFunc() string {
	return BaseFunc() + "->" + overrideVariable + "v1.Linux (from ostesting_override_linux.go)"
}

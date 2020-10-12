package main

var overrideVariable = "windows.alim.override"

func OverrideFunc() string {
	return BaseFunc() + "->" + overrideVariable + "v1.Windows (from ostesting_override_windows.go)"
}

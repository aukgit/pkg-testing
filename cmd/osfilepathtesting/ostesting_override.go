package main

var overrideVariable = "alim.override"

func OverrideFunc() string {
	return overrideVariable + BaseFunc() + "v1(from ostesting_override.go)"
}

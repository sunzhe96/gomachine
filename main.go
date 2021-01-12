package main

import (
	"github.com/sunzhe96/gomachine/stackvm"
)


func main() {
	vm := stackvm.Newvm()
	prog := []int{3, 4, 0x40000001, -100, 0x40000001, 0x40000000}
	vm.LoadProgram(prog)
	vm.Run()
}

package stackvm

/*
   Instruction format
   header: 2 bits
   data: 30 bits

   header format:
   0 => positive integer
   1 => primitive instruction
   2 => negative intger
   3 => undefined
*/

import "fmt"


type stackVM struct {
	pc int // program counter
	sp int // stack pointer
	memory []int
	typ int
	dat int
	running int
}

func Newvm() stackVM {
	var vm stackVM = stackVM {
		pc : 100,
		sp : 0,
		memory : make([]int, 1000, 1000),
		typ : 0,
		dat : 0,
		running : 1,
	}
	return vm
}

func (vm *stackVM) getType(instruction int) {
	var typ int = 0xc0000000
	vm.typ = (typ & instruction) >> 30
}

func (vm *stackVM) getData(instruction int) {
	var data int = 0x3fffffff
	vm.dat = data & instruction
}

func (vm *stackVM) fetch() {
	vm.pc++
}

func (vm *stackVM) decode() {
	vm.getType(vm.memory[vm.pc])
	vm.getData(vm.memory[vm.pc])
}

func (vm *stackVM) execute() {
	if vm.typ == 0 || vm.typ == 2 {
		vm.sp++
		vm.memory[vm.sp] = vm.dat
	} else {
		vm.doPrimitive()
	}
}

func (vm *stackVM) doPrimitive() {
	switch vm.dat {
	case 0: // halt
		fmt.Println("halt")
		vm.running = 0
		break
	case 1: // add
		fmt.Println("add ", vm.memory[vm.sp - 1], vm.memory[vm.sp])
		vm.memory[vm.sp - 1] = vm.memory[vm.sp - 1] + vm.memory[vm.sp]
		vm.sp--
		break
	}
}

// exported

func (vm *stackVM) Run() {
	vm.pc -= 1
	for vm.running == 1 {
		vm.fetch()
		vm.decode()
		vm.execute()
		fmt.Println("the top of the stack is: ", vm.memory[vm.sp])
	}
}

func (vm *stackVM) LoadProgram(prog []int) {
	for i, v := range prog {
		vm.memory[vm.pc + i] = v
	}
}

package pkg

import "fmt"

type AsusComputer struct {
	Memory int
	Cpu    int
}

func (pc AsusComputer) PrintDetails() {
	fmt.Printf("[Asus] pc CPU:[%d] Mem:[%d]\n", pc.Cpu, pc.Memory)
}

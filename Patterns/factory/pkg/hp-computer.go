package pkg

import "fmt"

type HpComputer struct {
	Memory int
	Cpu    int
}

func (pc HpComputer) PrintDetails() {
	fmt.Printf("[Hp] pc CPU:[%d] Mem:[%d]\n", pc.Cpu, pc.Memory)
}

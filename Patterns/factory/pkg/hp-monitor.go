package pkg

import "fmt"

type HpMonitor struct {
	Size int
}

func (m HpMonitor) PrintDetails() {
	fmt.Printf("[Hp] monitor size:[%d]", m.Size)
}

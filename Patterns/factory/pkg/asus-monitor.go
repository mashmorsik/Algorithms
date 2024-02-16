package pkg

import "fmt"

type AsusMonitor struct {
	Size int
}

func (m AsusMonitor) PrintDetails() {
	fmt.Printf("[Asus] monitor size:[%d]", m.Size)
}

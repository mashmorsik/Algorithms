package main

import (
	"Algorithms/Patterns/factory/pkg"
	"fmt"
)

var (
	brands = []string{pkg.Asus, pkg.HP, "dell"}
)

func main() {
	for _, brand := range brands {
		factory, err := pkg.GetFactory(brand)
		if err != nil {
			fmt.Println(err)
			continue
		}
		computer := factory.GetComputer()
		computer.PrintDetails()
		monitor := factory.GetMonitor()
		monitor.PrintDetails()
	}
}

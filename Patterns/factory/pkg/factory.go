package pkg

import (
	"errors"
	"fmt"
)

type Factory interface {
	GetComputer() Computer
	GetMonitor() Monitor
}

func GetFactory(brand string) (Factory, error) {
	switch brand {
	case Asus:
		return &AsusFactory{}, nil
	case HP:
		return &HpFactory{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Производитель %s не найден!", brand))
	}
}

package pkg

type AsusFactory struct {
}

func (factory AsusFactory) GetComputer() Computer {
	return AsusComputer{
		Memory: 8,
		Cpu:    4,
	}
}

func (factory AsusFactory) GetMonitor() Monitor {
	return AsusMonitor{
		Size: 15,
	}
}

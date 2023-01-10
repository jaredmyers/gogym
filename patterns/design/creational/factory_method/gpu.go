package factory

import "fmt"

// -- Abstract Factory for Making GPUs -- //
type factoryGpuAssembler interface {
	CreateGpu() gpuAssembler
}

func GetGpuAssembler(family string) (factoryGpuAssembler, error) {
	if family == "msi" {
		return &msiManufacturer{}, nil
	}
	if family == "asus" {
		return &asusManufacturer{}, nil
	}
	return nil, fmt.Errorf("Wrong type passed..")
}

// -- Conrete factory for MSI Family-- //
type msiManufacturer struct{}

func (m *msiManufacturer) CreateGpu() gpuAssembler {
	return &msiGpu{
		gpu{manufacturer: "msi"},
	}
}

// -- Concrete factory for ASUS Family -- //
type asusManufacturer struct{}

func (a *asusManufacturer) CreateGpu() gpuAssembler {
	return &asusGpu{
		gpu{manufacturer: "asus"},
	}
}

// -- Concrete Factory for MSI Family Products
type msiGpu struct {
	gpu
}

// -- Concrete Factory for ASUS Family Products -- //
type asusGpu struct {
	gpu
}

type gpuAssembler interface {
	Assemble(int, int)
	CheckAssembly() (int, int)
}

// -- Concrete Factory for Base product
type gpu struct {
	manufacturer string
	power        int
	speed        int
}

func (m *gpu) Assemble(power, speed int) {
	m.power = power
	m.speed = speed
}

func (m *gpu) CheckAssembly() (int, int) {
	return m.power, m.speed
}

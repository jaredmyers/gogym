package abstractfactory

import "fmt"

// --- Abstract Factory --- //
type furnitureCreator interface {
	CreateChair() chairCreator
	CreateCoffeeTable() coffeeTableCreator
	CreateSofa() sofaCreator
}

func GetFurnitureCreator(family string) (furnitureCreator, error) {
	if family == "ArtDeco" {
		return &artDeco{}, nil
	}
	if family == "Modern" {
		return &modern{}, nil
	}
	return nil, fmt.Errorf("Wrong Family Type Passed..")
}

// --- Concrete Factory for Art Deco Family --- //
type artDeco struct{}

func (A *artDeco) CreateChair() chairCreator {
	return &artDecoChair{
		chair: chair{logo: "ArtDeco", size: 12},
	}

}
func (A *artDeco) CreateCoffeeTable() coffeeTableCreator {
	return &artDecoCoffeeTable{
		coffeeTable: coffeeTable{logo: "ArtDeco", size: 12},
	}

}
func (A *artDeco) CreateSofa() sofaCreator {
	return &artDecoSofa{
		sofa: sofa{logo: "ArtDeco", size: 12},
	}

}

// --- Concrete Factory for Modern Family --- //
type modern struct{}

func (A *modern) CreateChair() chairCreator {
	return &modernChair{
		chair: chair{logo: "Modern", size: 100},
	}

}
func (A *modern) CreateCoffeeTable() coffeeTableCreator {
	return &modernCoffeeTable{
		coffeeTable: coffeeTable{logo: "Modern", size: 100},
	}

}
func (A *modern) CreateSofa() sofaCreator {
	return &modernSofa{
		sofa: sofa{logo: "Modern", size: 100},
	}

}

// --- Concrete Factory for Art Deco Products --- //
type artDecoChair struct {
	chair
}
type artDecoCoffeeTable struct {
	coffeeTable
}
type artDecoSofa struct {
	sofa
}

// --- Concrete Factory for Modern Products --- //
type modernChair struct {
	chair
}
type modernCoffeeTable struct {
	coffeeTable
}
type modernSofa struct {
	sofa
}

// --- Abstract Factory for Base Products -- //
type chairCreator interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}
type coffeeTableCreator interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}
type sofaCreator interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

// --- Concrete Factory for Base Products --- //
type chair struct {
	logo string
	size int
}

func (c *chair) SetLogo(logo string) {
	c.logo = logo
}
func (c *chair) GetLogo() string {
	return c.logo
}
func (c *chair) SetSize(size int) {
	c.size = size
}
func (c *chair) GetSize() int {
	return c.size
}

type coffeeTable struct {
	logo string
	size int
}

func (c *coffeeTable) SetLogo(logo string) {
	c.logo = logo
}
func (c *coffeeTable) GetLogo() string {
	return c.logo
}
func (c *coffeeTable) SetSize(size int) {
	c.size = size
}
func (c *coffeeTable) GetSize() int {
	return c.size
}

type sofa struct {
	logo string
	size int
}

func (s *sofa) SetLogo(logo string) {
	s.logo = logo
}
func (s *sofa) GetLogo() string {
	return s.logo
}
func (s *sofa) SetSize(size int) {
	s.size = size
}
func (s *sofa) GetSize() int {
	return s.size
}

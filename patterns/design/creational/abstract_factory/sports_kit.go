package abstractfactory

import "fmt"

// --- Abstract Factory --- //
type sportsFactory interface {
	MakeShirt() shirtMaker
	MakeShoe() shoeMaker
}

func GetSportsFactory(brand string) (sportsFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}
	if brand == "nike" {
		return &nike{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}

// --- Concrete Factory for Nike Family --- //
type nike struct{}

func (n *nike) MakeShirt() shirtMaker {
	return &nikeShirt{
		shirt: shirt{logo: "nike", size: 100},
	}
}
func (n *nike) MakeShoe() shoeMaker {
	return &nikeShoe{
		shoe: shoe{logo: "nike", size: 2},
	}
}

// --- Concrete Factory for Nike Products --- //
type nikeShirt struct {
	shirt
}
type nikeShoe struct {
	shoe
}

// --- Concrete Factory for Adidas Family --- //
type adidas struct{}

func (a *adidas) MakeShirt() shirtMaker {
	return &adidasShirt{
		shirt: shirt{logo: "adidas", size: 14},
	}
}
func (a *adidas) MakeShoe() shoeMaker {
	return &adidasShoe{
		shoe: shoe{logo: "adidas", size: 14},
	}
}

// --- Concrete Factory for Adidias Products --- //
type adidasShirt struct {
	shirt
}
type adidasShoe struct {
	shoe
}

// --- Abstract Product --- //
type shirtMaker interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

// --- Concrete Product --- //
type shirt struct {
	logo string
	size int
}

func (s *shirt) SetLogo(logo string) {
	s.logo = logo
}
func (s *shirt) GetLogo() string {
	return s.logo
}
func (s *shirt) SetSize(size int) {
	s.size = size
}
func (s *shirt) GetSize() int {
	return s.size
}

// --- Abstract Product --- //
type shoeMaker interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

// --- Concrete Product --- //
type shoe struct {
	logo string
	size int
}

func (s *shoe) SetLogo(logo string) {
	s.logo = logo
}
func (s *shoe) GetLogo() string {
	return s.logo
}
func (s *shoe) SetSize(size int) {
	s.size = size
}
func (s *shoe) GetSize() int {
	return s.size
}

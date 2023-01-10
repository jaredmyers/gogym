package factory

import "fmt"

// -- Simple Factory -- //
type GunCreator interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}

func CreateGun(gunType string) (GunCreator, error) {
	if gunType == "AK47" {
		return newAK(), nil
	}
	if gunType == "Musket" {
		return newMusket(), nil
	}

	return nil, fmt.Errorf("wrong type passed")
}

// -- Concrete HigherOrder Products --
type aK47 struct {
	gun
}

func newAK() GunCreator {
	return &aK47{
		gun: gun{name: "AK47", power: 7},
	}
}

type musket struct {
	gun
}

func newMusket() GunCreator {
	return &musket{
		gun: gun{name: "Musket", power: 1999},
	}
}

// -- Concrete Base Product -- //
type gun struct {
	name  string
	power int
}

func (g *gun) SetName(name string) {
	g.name = name
}
func (g *gun) GetName() string {
	return g.name
}

func (g *gun) SetPower(power int) {
	g.power = power
}
func (g *gun) GetPower() int {
	return g.power
}

package main

import "fmt"

// flyweight concrete factory
const (
	MercenaryDressType = "mDress"
	SpyDressType       = "sDress"
)

var dressFactorySingleInstance = &DressFactory{
	dressMap: make(map[string]Dress),
}

type DressFactory struct {
	dressMap map[string]Dress
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}

	if dressType == MercenaryDressType {
		d.dressMap[dressType] = newMercenaryDress("red")
		return d.dressMap[dressType], nil
	}

	if dressType == SpyDressType {
		d.dressMap[dressType] = newSpyDress("blue")
		return d.dressMap[dressType], nil
	}

	return nil, fmt.Errorf("wrong dress type passed")
}

func getDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}

// abstract dress
type Dress interface {
	getColor() string
}

// concrete flyweights
type MercenaryDress struct {
	color string
}

func (m *MercenaryDress) getColor() string {
	return m.color
}

func newMercenaryDress(color string) *MercenaryDress {
	return &MercenaryDress{color: color}
}

type SpyDress struct {
	color string
}

func (m *SpyDress) getColor() string {
	return m.color
}

func newSpyDress(color string) *SpyDress {
	return &SpyDress{color: color}
}

type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *Player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *Player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

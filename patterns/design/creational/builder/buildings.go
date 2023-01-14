package builder

import "fmt"

type builder interface {
	setWindows()
	setDoors()
	setFloors()
	getHouse() *house
}

func GetBuilder(builderType string) builder {
	if builderType == "normal" {
		return newNormalBuilder()
	}

	return nil
}

// Concrete For Normal Builder
type NormalBuilder struct {
	windows string
	doors   string
	floors  int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (n *NormalBuilder) setWindows() {
	n.windows = "Wood"
}
func (n *NormalBuilder) setDoors() {
	n.doors = "Wood"
}
func (n *NormalBuilder) setFloors() {
	n.floors = 2
}
func (n *NormalBuilder) getHouse() *house {
	return &house{
		doors:   n.doors,
		windows: n.windows,
		floors:  n.floors,
	}
}

// -- Concrete for base house

type house struct {
	windows string
	doors   string
	floors  int
}

func (h *house) GetInfo() {
	fmt.Println(h.windows)
	fmt.Println(h.doors)
	fmt.Println(h.floors)
}

// -- Director
type director struct {
	builder
}

func NewDirector(b builder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) SetBuilder(b builder) {
	d.builder = b
}

func (d *director) BuildHouse() *house {
	d.builder.setDoors()
	d.builder.setWindows()
	d.builder.setFloors()

	return d.builder.getHouse()
}

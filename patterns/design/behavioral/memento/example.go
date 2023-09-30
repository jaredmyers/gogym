package memento

type Originator struct {
	State1 string
	State2 int
	State3 bool
}

func (e *Originator) CreateMemento() *Memento {
	return &Memento{
		State1: e.State1,
		State2: e.State2,
		State3: e.State3,
	}
}

func (e *Originator) RestoreMemento(m *Memento) {
	if m != nil {
		e.State1, e.State2, e.State3 = m.GetSavedState()
	}
}

func (e *Originator) SetState(state1 string, state2 int, state3 bool) {
	e.State1 = state1
	e.State2 = state2
	e.State3 = state3
}

func (e *Originator) GetState() (string, int, bool) {
	return e.State1, e.State2, e.State3
}

type Memento struct {
	State1 string
	State2 int
	State3 bool
}

func (m *Memento) GetSavedState() (string, int, bool) {
	return m.State1, m.State2, m.State3
}

type Caretaker struct {
	MementoArray []*Memento
}

func (c *Caretaker) AddMemento(m *Memento) {
	c.MementoArray = append(c.MementoArray, m)
}

func (c *Caretaker) NextMemento() *Memento {
	if len(c.MementoArray) > 0 {
		next := c.MementoArray[len(c.MementoArray)-1]
		c.MementoArray = c.MementoArray[:len(c.MementoArray)-1]
		return next
	}

	return nil
}

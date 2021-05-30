package Memento

type Memento struct {
	value int
}

func NewMemento(value int) *Memento {
	return &Memento{
		value: value,
	}
}

type Number struct {
	value int
}

func NewNumber(value int) *Number {
	return &Number{
		value: value,
	}
}

func (th *Number) Double() {
	th.value *= 2
}

func (th *Number) Half() {
	th.value /= 2
}

func (th *Number) Value() int {
	return th.value
}

func (th *Number) CreateMemento() *Memento {
	return NewMemento(th.value)
}

// ReinstateMemento [ˌriːɪnˈsteɪt] 恢复备忘录
func (th *Number) ReinstateMemento(m *Memento) {
	th.value = m.value
}

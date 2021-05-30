package Iterator

type Iterator interface {
	Index() int
	Value() interface{}
	HasNext() bool
	Next()
}

type ArrayIterator struct {
	array []interface{}
	index int
}

func (th *ArrayIterator) Index() int {
	return th.index
}

func (th *ArrayIterator) Value() interface{} {
	return th.array[th.index]
}

func (th *ArrayIterator) HasNext() bool {
	return th.index+1 < len(th.array)
}

func (th *ArrayIterator) Next() {
	if th.HasNext() {
		th.index++
	}
}


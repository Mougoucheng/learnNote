package Command

import "testing"

func TestCommand_Execute(t *testing.T) {
	wan := NewPerson("wan", NewCommand(nil, nil))
	zhang := NewPerson("zhang", NewCommand(&wan, wan.Listen))
	fang := NewPerson("fang", NewCommand(&zhang, zhang.Buy))
	ding := NewPerson("ding", NewCommand(&fang, fang.Cook))
	li := NewPerson("li", NewCommand(&ding, ding.Wash))

	li.Talk()
}

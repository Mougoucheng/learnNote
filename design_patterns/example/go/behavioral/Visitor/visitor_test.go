package Visitor

import "testing"

func TestMe_Accept(t *testing.T) {
	me := &Me{}
	me.Accept(&Daddy{})
	me.Accept(&Mammy{})
	me.Accept(&Sister{})
	me.Accept(&Brother{})
	me.Accept(&Stranger{})
}

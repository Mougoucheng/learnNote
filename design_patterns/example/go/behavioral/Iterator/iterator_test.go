package Iterator

import (
	"log"
	"testing"
)

func TestArrayIterator(t *testing.T) {
	array := []interface{}{2, 6, 3, 4, 5, 1, 7}
	i := 0
	iterator := ArrayIterator{array: array, index: i}
	for it := iterator; it.HasNext(); it.Next() {
		idx, val := it.Index(), it.Value().(int)
		if val != array[idx] {
			log.Println("error")
		}
		log.Println(idx, val)
	}
}

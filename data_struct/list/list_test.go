package list

import (
	"testing"
)

func TestNewSingle(t *testing.T) {
	l := NewSingle()
	l.InsertHead(1)
	l.InsertHead(2)
	l.InsertHead(3)
	l.InsertHead(4)
	l.Print()
	l.ReverseV2()
	l.Print()
}

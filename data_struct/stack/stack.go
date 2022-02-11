package stack

import "errors"

// 20,155,232,844,224,682,496.
type ArrayStack struct {
	Values  []int
	current int
	length  int
}

func (a *ArrayStack) Push(value int) error {
	if a.current == a.length {
		return errors.New("no more")
	}
	a.Values[a.current] = value
	return nil
}

func (a *ArrayStack) Pop() (int, error) {
	if a.current == 0 {
		return 0, errors.New("no elem")
	}
	temp := a.Values[a.current]
	a.current--
	return temp, nil
}

func NewArrayStack(length int) *ArrayStack {
	return &ArrayStack{
		Values:  make([]int, length),
		current: 0,
		length:  length,
	}
}

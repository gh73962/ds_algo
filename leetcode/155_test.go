package leetcode

import (
	"log"
	"testing"
)

func TestConstructor(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		ms := Constructor()
		ms.Top()
		ms.Pop()
		ms.GetMin()
	})

	t.Run("2", func(t *testing.T) {
		ms := Constructor()
		ms.Push(3)
		log.Printf("ms.Top()=%d,min=%d", ms.Top(), ms.GetMin())
		ms.Push(5)
		log.Printf("ms.Top()=%d,min=%d", ms.Top(), ms.GetMin())
		ms.Push(2)
		log.Printf("ms.Top()=%d,min=%d", ms.Top(), ms.GetMin())

		ms.Pop()
		log.Printf("ms.GetMin()=%d", ms.GetMin())
		ms.Pop()
		log.Printf("ms.GetMin()=%d", ms.GetMin())
		ms.Pop()
		log.Printf("ms.GetMin()=%d", ms.GetMin())
		ms.Pop()
		log.Printf("ms.GetMin()=%d", ms.GetMin())
	})
}

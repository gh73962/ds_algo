package leetcode

// 请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）
// 你 只能 使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
// 你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。

type MyQueue struct {
	arr        []int
	head, tail int
}

func ConstructorMyQueue() MyQueue {
	return MyQueue{}
}

// 将元素 x 推到队列的末尾
func (this *MyQueue) Push(x int) {
	this.arr = append(this.arr, x)
	this.tail++
}

// 从队列的开头移除并返回元素
func (this *MyQueue) Pop() int {
	return 0
}

// 返回队列开头的元素
func (this *MyQueue) Peek() int {
	return 0
}

// 如果队列为空，返回 true ；否则，返回 false
func (this *MyQueue) Empty() bool {
	return false
}

package leetcode

// 输入：
// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]

// 输出：
// [null,null,null,null,-3,null,0,-2]

// 解释：
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin();   --> 返回 -3.
// minStack.pop();
// minStack.top();      --> 返回 0.
// minStack.getMin();   --> 返回 -2.

type MinStack struct {
	Stack  []int
	min    int
	length int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.length == 0 {
		this.min = val
	}
	this.Stack = append(this.Stack, val)
	this.length++
	if val < this.min {
		this.min = val
	}
}

func (this *MinStack) Pop() {
	switch this.length {
	case 0:
		return
	case 1:
		this.Stack = nil
		this.length = 0
		this.min = 0
	default:
		top := this.Top()
		this.Stack = this.Stack[:this.length-1]
		if top == this.min {
			for i, v := range this.Stack {
				if i == 0 {
					this.min = v
				}
				if this.min > v {
					this.min = v
				}
			}
		}
		this.length--
	}
}

func (this *MinStack) Top() int {
	switch this.length {
	case 0:
		return 0
	default:
		return this.Stack[this.length-1]
	}
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

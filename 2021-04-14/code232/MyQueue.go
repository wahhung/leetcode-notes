package main

// 232. 用栈实现队列
// 栈：后进先出
// 队列：先进先出
// 解题思路：双栈
type MyQueue struct {
	inStack []int
	outStack []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}


// 将元素 x 推到队列的末尾
// 将元素放到 inStack
func (this *MyQueue) Push(x int)  {
	this.inStack = append(this.inStack, x)
}

// 发生出列或者查询列时，将 inStack 搬到 ouStack
func (this *MyQueue) inToOut(){
	for len(this.inStack) > 0 {
		this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
		this.inStack = this.inStack[:len(this.inStack)-1]
	}
}


//从队列的开头移除并返回元素
// 出栈的数据没有，需要将入栈的数据搬过来
func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 {
		this.inToOut()
	}
	if len(this.outStack) == 0 {
		return -1 // 假设没有负数
	}
	val := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return val
}


// 返回队列开头的元素
func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		this.inToOut()
	}
	if len(this.outStack) == 0 {
		return -1 // 假设没有负数
	}
	return this.outStack[len(this.outStack)-1]
}


// 如果队列为空，返回 true ；否则，返回 false
func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

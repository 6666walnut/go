/**
* @Author: Chicken dishes
* @Date: 2021/5/27 20:20
**/

package main

import "fmt"

/**
 * 定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。
 * 示例:
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.min();   --> 返回 -3.
 * minStack.pop();
 * minStack.top();   --> 返回 0.
 * minStack.min();   --> 返回 -2.
 */

type MinStack struct {
	stack []int
	min   []int
}

/** initialize your data structure here. */
func Constructor1() MinStack {
	return MinStack{
		stack: []int{},
		min:   []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.min) == 0 || this.min[len(this.min)-1] >= x {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Pop() {
	end := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if end == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
func main() {
	obj := Constructor1()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.Min()
	fmt.Println(param_3, param_4)
}

/**
* @Author: Chicken dishes
* @Date: 2021/5/24 20:03
**/

package main

import (
	"container/list"
	"fmt"
)

/**
 * 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
 * 分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )
 */

type CQueue struct {
	stackIn  *list.List
	stackOut *list.List
}

func Constructor() CQueue {
	return CQueue{
		stackIn:  list.New(),
		stackOut: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stackIn.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	// 判断出栈长度是否为0
	if this.stackOut.Len() == 0 {
		for this.stackIn.Len() > 0 {
			this.stackOut.PushBack(this.stackIn.Remove(this.stackIn.Back()))
		}
	}
	if this.stackOut.Len() != 0 {
		e := this.stackOut.Back()
		this.stackOut.Remove(e)
		return e.Value.(int)
	}
	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
func main() {
	obj := Constructor()

	value := 2
	obj.AppendTail(value)
	param_2 := obj.DeleteHead()
	fmt.Println(param_2)
}

package dataStructure

type MyStack struct {
    queue1 []int
    queue2 []int

}


/** Initialize your data structure here. */
func Constructor() MyStack {
    return MyStack{
        queue1: make([]int, 0),
        queue2: make([]int, 0),
    }
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    if len(this.queue2) > 0 {
        for i := len(this.queue2)-1; i >= 0; i-- {
            this.queue1 = append(this.queue1, this.queue2[i])
        }
        this.queue2 = []int{}
    }
    this.queue1 = append(this.queue1, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    if len(this.queue2) > 0 {
        for i := len(this.queue2)-1; i >= 0; i-- {
            this.queue1 = append(this.queue1, this.queue2[i])
        }
        this.queue2 = []int{}
    }

    value := this.queue1[len(this.queue1)-1]
    this.queue1 = this.queue1[:len(this.queue1)-1]
    return value
}


/** Get the top element. */
func (this *MyStack) Top() int {
    if len(this.queue2) > 0 {
        return this.queue1[0]
    } else {
        return this.queue1[len(this.queue1)-1]
    }
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    return len(this.queue1) == 0 && len(this.queue2) == 0
}

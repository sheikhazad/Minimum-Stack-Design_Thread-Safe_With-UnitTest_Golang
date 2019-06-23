/************************************************************************
* This is a thread safe and exception safe Stack and MinStack.
* Non-Thread safe Stack is implemented in other repository:
* https://github.com/sheikhazad/MyStack-No_Thread_Safe-Golang
*************************************************************************/
package minimumStack

import (
	"errors"
	"fmt"
	"sync"
)

type MinStack struct {
	stack    []int //Normal stack
	minStack []int //Stack to track min values
	mutex    sync.Mutex
}

func Constructor() *MinStack {
	return &MinStack{make([]int, 0), make([]int, 0), sync.Mutex{}}
}

//Push data into the stack
func (this *MinStack) Push(element int) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	//Push into minStack
	minSize := len(this.minStack)
	if minSize == 0 || element <= this.minStack[minSize-1] {
		this.minStack = append(this.minStack, element)
	}

	this.stack = append(this.stack, element)
}

/*
  1. Pop last/top element from the stack
  3. Return last/top element from stack before popping
*/
func (this *MinStack) Pop() {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	size := len(this.stack)
	if size == 0 {
		fmt.Println("Stack is empty.")
		return
	}

	minSize := len(this.minStack)

	//min stack is popped ONLY if the top values (BEFORE popping) of the min stack and normal stack are equal.
	if this.stack[size-1] == this.minStack[minSize-1] {
		this.minStack = this.minStack[:minSize-1]
	}

	//Remove top element
	this.stack = (this.stack)[:size-1]
}

func (this MinStack) Top() (int, error) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	size := len(this.stack)
	if size == 0 {
		return 0, errors.New("Stack is empty.")
	}

	return this.stack[size-1], nil
}

func (this MinStack) getMin() (int, error) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	minSize := len(this.minStack)
	if minSize == 0 {
		return 0, errors.New("Min Stack is empty.")
	}
	return this.minStack[minSize-1], nil
}

func (this MinStack) isEmpty() bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	return len(this.stack) == 0
}

func (this *MinStack) Clear() {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.stack = nil
	this.minStack = nil
}

func (this MinStack) Size() int {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	return len(this.stack)
}

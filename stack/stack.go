package stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxSize int
	Array   [5]int // 切片
	Top     int    // 栈顶指针
}

// IsEmpty 栈空
func (stack *Stack) IsEmpty() bool {
	return stack.Top == -1
}

// IsFull 栈满
func (stack *Stack) IsFull() bool {
	return stack.Top == stack.MaxSize-1
}

// Push 入栈
func (stack *Stack) Push(val int) (err error) {
	if stack.IsFull() {
		return errors.New("the stack is full")
	}
	stack.Top++
	stack.Array[stack.Top] = val
	return nil
}

// Pop 出栈
func (stack *Stack) Pop() (val int, err error) {
	if stack.IsEmpty() {
		return -1, errors.New("the stack is empty")
	}
	val = stack.Array[stack.Top]
	stack.Top--
	return val, nil
}

// TraversalStack 遍历栈元素
func (stack *Stack) Traversal() {
	if stack.IsEmpty() {
		fmt.Println("the stack is empty")
		return
	}
	for i := stack.Top; i >= 0; i-- {
		fmt.Printf("Stack[%d]: %d; \t", stack.Top-i, stack.Array[i])
	}
}

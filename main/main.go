package main

import (
	"data_structures_and_algorithms/cycleQueue"
	"data_structures_and_algorithms/queue"
	"data_structures_and_algorithms/stack"
	"fmt"
)

func main() {
	// Queue
	fmt.Println("===== Queue =====")
	myQueue := &queue.Queue{
		MaxSize: 5,
		Front:   -1,
		Rear:    -1,
	}
	fmt.Print("初始化-打印队列元素：")
	myQueue.TraversalQueue()
	// 入队
	for i := 0; i < 4; i++ {
		err := myQueue.AddQueue(i)
		if err != nil {
			return
		}
	}
	fmt.Println()
	fmt.Print("添加元素-打印队列元素：")
	myQueue.TraversalQueue()
	// 获取
	index := 1
	val, err := myQueue.GetEleByQueue(index)
	if err != nil {
		return
	}
	fmt.Println()
	fmt.Printf("获取下标为%d的元素：%d", index, val)
	// 出队
	shiftQueue, err := myQueue.ShiftQueue()
	if err != nil {
		return
	}
	fmt.Println()
	fmt.Printf("当前出队元素：%d", shiftQueue)
	fmt.Println()
	fmt.Print("出队-打印队列元素")
	myQueue.TraversalQueue()
	fmt.Println()

	// CycleQueue
	fmt.Println("===== CycleQueue =====")
	myCycleQueue := &cycleQueue.CycleQueue{
		MaxSize: 5,
		Front:   0,
		Rear:    0,
	}
	// 入队
	for i := 1; i <= 4; i++ {
		err = myCycleQueue.EnQueue(i)
		if err != nil {
			return
		}
	}
	fmt.Println("添加元素：", myCycleQueue.Array)
	err = myCycleQueue.EnQueue(5)
	if err != nil {
		fmt.Println("错误信息：", err)
		//return
	}
	fmt.Println("队列元素：", myCycleQueue.Array)
	fmt.Println("打印队首元素：", myCycleQueue.GetHead())
	deQueue, err := myCycleQueue.DeQueue()
	if err != nil {
		fmt.Println("出队的错误信息：", err)
	}
	fmt.Println("出队元素：", deQueue)
	fmt.Print("出队后的队列：")
	myCycleQueue.TraversalQueue()
	err = myCycleQueue.EnQueue(5)
	if err != nil {
		fmt.Println("错误信息：", err)
		//return
	}
	fmt.Print("添加元素：")
	myCycleQueue.TraversalQueue()
	for i := 0; i < 3; i++ {
		deQueue, err := myCycleQueue.DeQueue()
		if err != nil {
			fmt.Println("出队的错误信息：", err)
		}
		fmt.Println("出队元素：", deQueue)
	}
	fmt.Print("出队后的队列：")
	myCycleQueue.TraversalQueue()

	// Stack
	fmt.Println("===== Stack =====")
	myStack := &stack.Stack{
		MaxSize: 5,
		Top:     -1,
	}
	fmt.Print("Push前的Stack：")
	myStack.Traversal()
	for i := 0; i < 5; i++ {
		err = myStack.Push(i + 1)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Print("Push后的Stack：")
	myStack.Traversal()
	fmt.Println()
	err = myStack.Push(7)
	if err != nil {
		fmt.Println(err)
	}
	stackHead, err := myStack.Pop()
	fmt.Println("StackHead: ", stackHead)
	if err != nil {
		fmt.Println(err)
		//return
	}
	fmt.Println()
	myStack.Traversal()
	fmt.Println()
	for i := 0; i < 5; i++ {
		stackHead, err := myStack.Pop()
		if err != nil {
			fmt.Println(err)
			//return
		}
		fmt.Println("StackHead: ", stackHead)
	}
	myStack.Traversal()
}

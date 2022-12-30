package main

import (
	"data_structures_and_algorithms/queue"
	"fmt"
)

func main() {
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
}

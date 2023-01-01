package cycleQueue

import (
	"errors"
	"fmt"
)

type CycleQueue struct {
	MaxSize int
	Array   [5]int
	Front   int
	Rear    int
}

// QueueFull 队满
func (queue *CycleQueue) QueueFull() bool {
	// 留一个空间
	return (queue.Rear+1)%queue.MaxSize == queue.Front
}

// QueueEmpty 队空
func (queue *CycleQueue) QueueEmpty() bool {
	// 首尾指针相同时候，队列为空
	return queue.Rear == queue.Front
}

// GetHead 返回队头元素
func (queue *CycleQueue) GetHead() int {
	return queue.Array[queue.Front]
}

// EnQueue 入队
func (queue *CycleQueue) EnQueue(value int) (err error) {
	if queue.QueueFull() {
		return errors.New("the queue is full")
	}
	queue.Array[queue.Rear] = value
	queue.Rear = (queue.Rear + 1) % queue.MaxSize
	return nil
}

// DeQueue 出队
func (queue *CycleQueue) DeQueue() (val int, err error) {
	if queue.QueueEmpty() {
		return -1, errors.New("the queue is empty")
	}
	val = queue.Array[queue.Front]
	queue.Front = (queue.Front + 1) % queue.MaxSize
	return val, nil
}

// QueueLength 队列长度
func (queue *CycleQueue) QueueLength() int {
	return (queue.Rear - queue.Front + queue.MaxSize) % queue.MaxSize
}

// TraversalQueue 遍历队列元素
func (queue *CycleQueue) TraversalQueue() {
	if queue.QueueLength() == 0 {
		fmt.Println("the queue is empty")
		return
	}
	tempFront := queue.Front
	for i := 0; i < queue.QueueLength(); i++ {
		fmt.Printf("arr[%d]=%d;\t", tempFront, queue.Array[tempFront])
		tempFront = (tempFront + 1) % queue.MaxSize
	}
	fmt.Println()
}

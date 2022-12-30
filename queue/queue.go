package queue

import (
	"errors"
	"fmt"
)

type Queue struct {
	MaxSize int
	Array   [5]int // 数组
	Front   int    // 队首 - 指向含有元素的前一个下标
	Rear    int    // 队尾 - 指向当前队尾元素的下标
}

// AddQueue - 入队操作
func (queue *Queue) AddQueue(val int) (err error) {
	if queue.Rear == queue.MaxSize-1 {
		// 队满了
		return errors.New("the queue is full")
	}
	queue.Rear++
	queue.Array[queue.Rear] = val
	return
}

// ShiftQueue - 出队操作
func (queue *Queue) ShiftQueue() (val int, err error) {
	if queue.Front == queue.Rear {
		// 队空了
		return -1, errors.New("the queue is empty")
	}
	queue.Front++
	val = queue.Array[queue.Front]
	return val, err
}

// TraversalQueue - 遍历队列
func (queue *Queue) TraversalQueue() {
	// 遍历从队首到队尾的所有元素
	for i := queue.Front + 1; i <= queue.Rear; i++ {
		fmt.Printf("Queue[%v]: %v;\t", i, queue.Array[i])
	}
}

// GetEleByQueue - 获取队列某个元素
func (queue *Queue) GetEleByQueue(index int) (val int, err error) {
	if !(index >= queue.Front && index <= queue.Rear) {
		return -1, errors.New("cannot get a value")
	}
	return queue.Array[index], err
}

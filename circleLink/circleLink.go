package circleLink

import (
	"errors"
	"fmt"
)

type CircleLink struct {
	Var  int
	Next *CircleLink
}

// InsertNode 添加元素
func (circleLink *CircleLink) InsertNode(val int) {
	if circleLink.Next == nil {
		circleLink.Next = &CircleLink{
			Var:  val,
			Next: circleLink,
		}
		return
	}

	temp := circleLink
	for {
		if temp.Next == circleLink {
			break
		}
		temp = temp.Next
	}
	temp.Next = &CircleLink{
		val,
		circleLink,
	}
}

// DeNode 删除指定元素
func (circleLink *CircleLink) DeNode(index int) (err error) {
	if index > circleLink.Length() || index < 0 {
		return errors.New("the index exceeds the actual length")
	}
	temp := circleLink
	for {
		if index <= 0 {
			fmt.Println("找到需要删除元素的前一个了")
			break
		}
		index--
		temp = temp.Next
	}
	temp.Next = temp.Next.Next
	return nil
}

// Length 返回长度
func (circleLink *CircleLink) Length() (length int) {
	temp := circleLink
	index := 0
	for {
		// 两个临界值，一个头结点不算，一个循环完成后
		if temp.Next == circleLink || temp.Next == nil {
			break
		}
		index++
		temp = temp.Next
	}
	return index
}

// TraversalCircleLink 遍历循环链表
func (circleLink *CircleLink) TraversalCircleLink() {
	temp := circleLink
	fmt.Println("循环遍历元素：")
	for {
		if temp.Next == circleLink || temp.Next == nil {
			fmt.Println("已经循环一遍了")
			break
		}
		fmt.Println(temp.Next.Var)
		temp = temp.Next
	}
}

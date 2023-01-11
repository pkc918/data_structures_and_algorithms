package doublyLink

import (
	"errors"
	"fmt"
)

type DoublyLink struct {
	Pre  *DoublyLink
	Next *DoublyLink
	Val  int
}

// ListPush 插入节点， 值为 val
func (linkList *DoublyLink) ListPush(val int) {
	temp := linkList
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	temp.Next = &DoublyLink{
		Val:  val,
		Next: nil,
		Pre:  temp.Next,
	}
}

// ListInsert 中间插入元素 第i个， 值为 val
func (linkList *DoublyLink) ListInsert(i int, val int) (err error) {
	if i > linkList.ListLength() || i < 0 {
		return errors.New("the index exceeds the length")
	}
	temp := linkList
	for i > 0 {
		i--
		temp = temp.Next
	}
	temp.Next = &DoublyLink{Val: val, Next: temp.Next, Pre: temp.Next.Pre}
	return nil
}

// ListDelete 中间删除元素 第i个
func (linkList *DoublyLink) ListDelete(i int) (err error) {
	if i > linkList.ListLength() || i < 0 {
		return errors.New("the index exceeds the length")
	}
	temp := linkList
	for i > 0 {
		i--
		temp = temp.Next
	}
	temp.Next = temp.Next.Next
	temp.Next.Pre = temp
	return nil
}

// ListLength 获取链表长度
func (linkList *DoublyLink) ListLength() (length int) {
	temp := linkList
	length = 0
	for {
		temp = temp.Next
		length++
		if temp.Next == nil {
			break
		}
	}
	return length
}

// TraversalDoublyLink 正序遍历
func (doublyLink *DoublyLink) TraversalDoublyLink() {
	temp := doublyLink
	index := 1
	for {
		fmt.Printf("第%d个元素：%d", index, temp.Next.Val)
		fmt.Println()
		temp = temp.Next
		index++
		if temp.Next == nil {
			break
		}
	}
}

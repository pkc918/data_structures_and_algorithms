package linkList

import (
	"errors"
	"fmt"
)

type LinkList struct {
	Val  int
	Next *LinkList
}

// ListPush 插入节点， 值为 val
func (linkList *LinkList) ListPush(val int) {
	temp := linkList
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	temp.Next = &LinkList{
		Val:  val,
		Next: nil,
	}
}

// ListInsert 中间插入元素 第i个， 值为 val
func (linkList *LinkList) ListInsert(i int, val int) (err error) {
	if i > linkList.ListLength() || i < 0 {
		return errors.New("the index exceeds the length")
	}
	temp := linkList
	for i > 0 {
		i--
		temp = temp.Next
	}
	temp.Next = &LinkList{Val: val, Next: temp.Next}
	return nil
}

// ListDelete 中间删除元素 第i个
func (linkList *LinkList) ListDelete(i int) (err error) {
	if i > linkList.ListLength() || i < 0 {
		return errors.New("the index exceeds the length")
	}
	temp := linkList
	for i > 0 {
		i--
		temp = temp.Next
	}
	temp.Next = temp.Next.Next
	return nil
}

// ListLength 获取链表长度
func (linkList *LinkList) ListLength() (length int) {
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

// TraversalLink 打印链表
func (linkList *LinkList) TraversalLink() {
	temp := linkList
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

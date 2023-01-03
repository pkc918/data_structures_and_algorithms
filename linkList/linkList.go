package linkList

import "fmt"

type LinkList struct {
	Val  int
	Next *LinkList
}

// ListInsert 插入 head，第i个，值为 val
func (linkList *LinkList) ListInsert(val int) {
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

// 打印链表
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

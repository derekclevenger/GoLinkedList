package main

import (
	"fmt"
)

type Node struct {
	value int
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (list *LinkedList) Append(v int) {
	node := Node{v, nil}
	if list.head == nil {
		list.head = &node
	} else {
		end := list.head
		for {
			if end.next == nil {
				break
			}
			end = end.next
		}
		end.next = &node
	}
	list.size += 1
}

func(list *LinkedList) InsertAt(ind int, v int) error {

	if ind < 0 || ind > list.size {
		return fmt.Errorf("Index out of bounds")
	}
	node := Node{v,nil}
	if ind == 0 {
		node.next = list.head
		list.head = &node
		return nil
	}
	ll := list.head
	c := 0
	for c < ind - 1 {
		c++
		ll = ll.next
	}
	node.next = ll.next
	ll.next = &node
	list.size += 1
	return nil
}

func (list *LinkedList) RemoveAt(i int) (*Node, error) {
	if i > list.size || i < 0 {
		return nil, fmt.Errorf("Index out of range")
	}
	j := 0
	node := list.head
	for j < i - 1{
		j++
		node = node.next
	}
	remove := node.next
	node.next = remove.next
	list.size -= 1

	return remove, nil
}
func (list *LinkedList) IndexOf(i int) int {

	node := list.head
	j := 0
	for {
		if j == i {
			return node.value
		}
		if node.next == nil {
			return -1
		}
		node = node.next
		j++
	}
}

func (list *LinkedList) IsEmpty() bool {
	return list.size == 0
}

func (list *LinkedList) Size() int {
	size := 1
	last := list.head
	for {
		if last == nil || last.next == nil {
			break
		}
		last = last.next
		size++
	}
	return size
}
func (list *LinkedList) PrintList() {
	head := list.head
	for{
		if head == nil {
			break
		}
		fmt.Println(head.value)
		head = head.next
	}
}

func main() {
	linkedList := LinkedList{}
	linkedList.Append(2)
	fmt.Printf("item at index 0 %d\n",linkedList.IndexOf(0))
	fmt.Println(linkedList.IsEmpty())
	linkedList.InsertAt(0,5)
	fmt.Printf("item at index 0 %d\n",linkedList.IndexOf(0))
	linkedList.Append(22)
	fmt.Printf("size is %d \n", linkedList.Size())
	linkedList.RemoveAt(1)
	fmt.Println(linkedList.Size())
	fmt.Printf("item at index 1 %d\n",linkedList.IndexOf(1))
	fmt.Printf("size is %d \n",linkedList.Size())
	linkedList.RemoveAt(1)
	fmt.Printf("size is %d \n",linkedList.Size())
	linkedList.Append(44)
	linkedList.Append(33)
	linkedList.PrintList()
	fmt.Printf("size is %d \n",linkedList.Size())
}
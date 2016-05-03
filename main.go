package main

import "fmt"

type node struct {
	data int
	prev *node
	next *node
}

type doublyLinkedList struct {
	head   *node
	tail   *node
	length int
}

func (l *doublyLinkedList) add(value int) {

	var tempNode = new(node)
	tempNode.data = value

	if l.head == nil { //first element
		l.head = tempNode
		l.tail = tempNode
	} else if l.head.next == nil {
		l.head.next = tempNode
		l.tail = tempNode
		l.tail.prev = l.head
	} else {
		tempNode.prev = l.tail
		l.tail.next = tempNode
		l.tail = tempNode
	}
	l.length++
}

func (l *doublyLinkedList) get(id int) (e bool, data int) {
	iterator := 0
	var head = l.head
	for head != nil {
		if iterator == id {

			return true, head.data
		}
		head = head.next
		iterator++
	}

	return false, 0
}

func (l *doublyLinkedList) show() {

	var x = l.head

	if x == nil {
		fmt.Println("List is empty")
		return
	}

	if x.next == nil {
		fmt.Printf("%d\n", x.data)
		return
	}

	for x != nil {
		fmt.Printf("%d\n", x.data)
		x = x.next
	}
	return
}

func main() {

	var list = new(doublyLinkedList)

	list.add(10)
	list.add(20)
	list.add(30)
	list.add(40)

	var e, d = list.get(11)

	if e {
		fmt.Println(d)
	} else {
		fmt.Printf("%s \n", "Sequence")
	}

	//fmt.Println(list.head.next.next.data)

}
 
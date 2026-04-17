package main

import "fmt"

// Single node of a Linked List
type Node struct {
	data int
	next *Node
}

// The Linked List
type LinkedList struct {
	start *Node
	end   *Node
}

func (l *LinkedList) addElement(data int) {
	temp := &Node{data: data}

	if l.start == nil {
		l.start = temp
	} else {
		l.end.next = temp
	}

	l.end = temp
}

func (l *LinkedList) removeElement(data int) bool {
	if l.start == nil {
		return false
	}

	if l.start.data == data {
		l.start = l.start.next

		if l.start == nil {
			l.end = nil
		}

		return true
	}

	curr := l.start.next
	prevc := l.start

	for curr != nil {
		if curr.data == data {
			prevc.next = curr.next

			if curr == l.end {
				l.end = prevc
			}

			return true
		}
		prevc = curr
		curr = curr.next
	}
	return false
}

func (l *LinkedList) showElements() {
	ll := l.start

	fmt.Println("Linked list:")

	for ll != nil {
		fmt.Println(ll.data)
		ll = ll.next
	}
}

func main() {

	var c1 string
	var ci int

	var l LinkedList

	fmt.Println("Use 'see' command to list all commands")

	for {
		fmt.Print(">> ")
		fmt.Scan(&c1)

		if c1 == "exit" {
			break
		}

		switch c1 {
		case "add":
			fmt.Scan(&ci)
			l.addElement(ci)
			fmt.Println("Added element: ", ci)
		case "remove":
			fmt.Scan(&ci)
			if l.removeElement(ci) {
				fmt.Println("Element Removed: ", ci)
			} else {
				fmt.Println("Element not found")
			}
		case "show":
			l.showElements()

		case "see":
			fmt.Println("add <int>")
			fmt.Println("remove <int>")
			fmt.Println("show")
			fmt.Println("exit")
		default:
			fmt.Println("Invalid command")
		}

	}
}

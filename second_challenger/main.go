package main

import "fmt"

//SinglyLinkedListNode structure of objects
type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func removeNodes(listHead *SinglyLinkedListNode, x int32) *SinglyLinkedListNode {
	okays := []int32{}
	response := &SinglyLinkedListNode{}

	for {
		if listHead == nil {
			break
		}

		//Searches for numbers less than X
		if listHead.data <= x {
			okays = append(okays, listHead.data)
		}
		listHead = listHead.next
	}

	//No number within the standards
	if len(okays) == 0 {
		return nil
	}

	//One or more number within the standards
	if len(okays) >= 1 {
		response.data = okays[0]
		//Many numbers within the standards
		if len(okays) > 1 {
			addNext(response, okays[1:])
		}
	}
	return response
}

func show(current *SinglyLinkedListNode) {
	for {
		if current == nil {
			break
		}
		fmt.Println(current.data)
		current = current.next
	}
	fmt.Println("--- --- --- --- --- --- ---")
}

func addNext(item *SinglyLinkedListNode, itens []int32) *SinglyLinkedListNode {
	if len(itens) > 0 {
		newItem := &SinglyLinkedListNode{data: itens[0]}
		item.next = addNext(newItem, itens[1:])
	}
	return item
}

func main() {
	items := SinglyLinkedListNode{data: 1}
	addNext(&items, []int32{2, 3, 4, 5})
	show(&items)

	solved := removeNodes(&items, 3)
	show(solved)
}

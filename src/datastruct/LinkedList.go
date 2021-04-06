package datastruct

import "errors"

type LinkedList struct {
	size int
	head *Node
}

func (list *LinkedList) Append(add *Item) {
	addNode := makeNode(add, nil)
	//if list is empty, set add as the head
	if list.IsEmpty() {
		list.head = addNode
	} else {
		//find last item (aka tail) of linked list
		curr := list.head
		for curr.next != nil {
			curr = curr.next
		}
		//set tail.next to the node
		curr.next = addNode
	}
	list.size++
}

func (list *LinkedList) InsertAfter(i int, target *Item) error {
	//empty check
	if list.IsEmpty() {
		return errors.New("list is empty")
	}
	//param out of bounds check
	if i > list.size || i <= 0 {
		return errors.New("param index out of bounds")
	}
	//if the target is the head, then insert after the head
	if i == 1 {
		temp := makeNode(target, list.head.next)
		list.head.next = temp
		return nil
	} else {
		curr, index := list.head, 0
		for index <= list.size {
			if index == i {
				temp := makeNode(target, curr.next.next)
				curr.next = temp
				return nil
			}
			curr = curr.next
			index++
		}
	}
	return errors.New("index not found")
}

func (list *LinkedList) RemoveItem(target *Item) error {
	//empty check
	if list.IsEmpty() {
		return errors.New("list is empty")
	}
	//if head.value = target, set head to head.next
	if list.head.value == target {
		list.head = list.head.next
		return nil
	} else {
		//find the node before target
		curr := list.head
		for curr.next != nil {
			if curr.next.value == target {
				temp := curr.next.next
				curr.next.next = nil
				curr.next = temp
				return nil
			}
			curr = curr.next
		}
	}
	return errors.New("Item not found")
}

func (list *LinkedList) RemoveIndex(i int) (*Item, error) {
	if list.IsEmpty() {
		return nil, errors.New("list is empty")
	}
	if i > list.size || i <= 0 {
		return nil, errors.New("param index out of bounds")
	}
	//if i == 1, then remove head by setting head to head.next
	if i == 1 {
		item := list.head.value
		list.head = list.head.next
		return &item, nil
	} else {
		//find the node according to the given index
		curr, index := list.head, 0
		for index <= list.size {
			if index == i {
				val := curr.value
				temp := curr.next.next
				curr.next.next = nil
				curr.next = temp
				return &val, nil
			}
			curr = curr.next
			index++
		}
	}
	return nil, nil
}

func (list *LinkedList) IsEmpty() bool {
	return list.head == nil
}

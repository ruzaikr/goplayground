package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type linkedListNode struct {
	value int
	next *linkedListNode
}

type MyLinkedList struct {
	head *linkedListNode
	tail *linkedListNode
	length int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		head: nil,
		tail: nil,
		length: 0,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (mll *MyLinkedList) Get(index int) int {
	if index < 0 || index > 1000 {
		return -1
	}
	var current = mll.head
	var i = 0
	for current != nil {
		if index == i {
			return current.value
		}
		current = current.next
		i++
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (mll *MyLinkedList) AddAtHead(val int)  {
	var n = linkedListNode{
		value: val,
		next:  mll.head,
	}
	mll.head = &n

	if mll.tail == nil {
		mll.tail = &n
	}

	mll.length++
}

/** Append a node of value val to the last element of the linked list. */
func (mll *MyLinkedList) AddAtTail(val int)  {
	var n = linkedListNode{
		value: val,
		next:  nil,
	}

	if mll.tail != nil {
		mll.tail.next = &n
	}
	mll.tail = &n

	if mll.head == nil {
		mll.head = &n
	}

	mll.length++
}

func (mll *MyLinkedList) Reset() {
	mll.head = nil
	mll.tail = nil
	mll.length = 0
}

func (mll *MyLinkedList) DeleteHead() {
	if mll.length < 1 {
		return
	}

	if mll.length == 1 {
		mll.Reset()
		return
	}

	mll.head = mll.head.next
	mll.length--
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (mll *MyLinkedList) AddAtIndex(index int, val int)  {
	if index < 0 || index > 1000 {
		return
	}

	if index > mll.length {
		return
	} else if index == mll.length {
		mll.AddAtTail(val)
		return
	} else if index == 0 {
		mll.AddAtHead(val)
	}

	var prev = mll.head
	var current = mll.head.next
	var i = 1
	for current != nil {
		if index == i {
			var n = linkedListNode{
				value: val,
				next:  current,
			}
			prev.next = &n
			mll.length++
			return
		}
		prev = current
		current = current.next
		i++
	}
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (mll *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 || index > 1000 || index >= mll.length {
		return
	}

	if index == 0 {
		mll.DeleteHead()
		return
	}

	var prev = mll.head
	var current = mll.head.next
	var i = 1
	for current != nil {
		if index == i {
			prev.next = current.next
			mll.length--
			if current.next == nil {
				mll.tail = prev
			}
			return
		}
		prev = current
		current = current.next
		i++
	}
}

func (mll *MyLinkedList) ToString()  {
	//fmt.Println("====BEGINNING====")

	if mll.length == 0 {
		fmt.Println("[]")
	}else {
		var current = mll.head

		var i = 0
		for current != nil {
			fmt.Print(i, ":[", current.value ,"]-->")
			current = current.next
			i++
		}
	}
	fmt.Println("\n")
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
	var ll = Constructor()

	inputFile, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	var scanner = bufio.NewScanner(inputFile)
	scanner.Scan()
	var fnNamesStr = scanner.Text()
	scanner.Scan()
	var fnParamsStr = scanner.Text()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var fnNamesStrNoQuotes = strings.Replace(fnNamesStr[16:len(fnNamesStr)-1], "\"", "", -1)
	var fnNames = strings.Split(fnNamesStrNoQuotes, ",")

	var fnParams = strings.Split(fnParamsStr[5:len(fnParamsStr)-2], "],[")

	for i, fnName := range fnNames {
		switch fnName {
		case "get":
			var fnParamInt, err = strconv.Atoi(fnParams[i])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Get", fnParamInt, ": ", ll.Get(fnParamInt), "\n")
		case "addAtHead":
			var fnParamInt, err = strconv.Atoi(fnParams[i])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("addAtHead(",fnParamInt,")")
			ll.AddAtHead(fnParamInt)
			ll.ToString()
		case "addAtTail":
			var fnParamInt, err = strconv.Atoi(fnParams[i])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("addAtTail(",fnParamInt,")")
			ll.AddAtTail(fnParamInt)
			ll.ToString()
		case "addAtIndex":
			var indexAndVal = strings.Split(fnParams[i], ",")
			var index, err = strconv.Atoi(indexAndVal[0])
			if err != nil {
				log.Fatal(err)
			}
			var val, err1 = strconv.Atoi(indexAndVal[1])
			if err1 != nil {
				log.Fatal(err1)
			}
			fmt.Println("addAtIndex(",index,",",val,")")
			ll.AddAtIndex(index, val)
			ll.ToString()
		case "deleteAtIndex":
			var fnParamInt, err = strconv.Atoi(fnParams[i])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("deleteAtIndex(",fnParamInt,")")
			ll.DeleteAtIndex(fnParamInt)
			ll.ToString()
		}
	}
}
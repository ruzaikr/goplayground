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
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index > 1000 {
		return -1
	}
	var current = this.head
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
func (this *MyLinkedList) AddAtHead(val int)  {
	var n = linkedListNode{
		value: val,
		next:  this.head,
	}
	this.head = &n

	if this.tail == nil {
		this.tail = &n
	}

	this.length++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	var n = linkedListNode{
		value: val,
		next:  nil,
	}
	if this.tail != nil {
		this.tail.next = &n
	}
	this.tail = &n

	if this.head == nil {
		this.head = &n
	}

	this.length++
}

func (this *MyLinkedList) Reset() {
	this.head = nil
	this.tail = nil
	this.length = 0
}

func (this *MyLinkedList) DeleteHead() {
	if this.length < 1 {
		return
	}

	if this.length == 1 {
		this.Reset()
		return
	}

	this.head = this.head.next
	this.length--
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index < 0 || index > 1000 {
		return
	}

	if index > this.length {
		return
	} else if index == this.length {
		this.AddAtTail(val)
		return
	} else if index == 0 {
		this.AddAtHead(val)
	}

	var prev = this.head
	var current = this.head.next
	var i = 1
	for current != nil {
		if index == i {
			var n = linkedListNode{
				value: val,
				next:  current,
			}
			prev.next = &n
			this.length++
			return
		}
		prev = current
		current = current.next
		i++
	}
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 || index > 1000 || index >= this.length {
		return
	}

	if index == 0 {
		this.DeleteHead()
		return
	}

	var prev = this.head
	var current = this.head.next
	var i = 1
	for current != nil {
		if index == i {
			prev.next = current.next
			return
		}
		prev = current
		current = current.next
		i++
	}
}

func (this *MyLinkedList) ToString()  {
	//fmt.Println("====BEGINNING====")

	if this.length == 0 {
		fmt.Println("[]")
	}else {
		var current = this.head

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

	for i, fnParam := range fnParams {
		fmt.Println(fnNames[i], " ", fnParam)
	}
	log.Fatal("hello")

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
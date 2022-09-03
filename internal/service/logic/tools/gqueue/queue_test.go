package gqueue

import (
	"fmt"
	"testing"
)

func TestQueueNode_GetPathString(t *testing.T) {
	path := "set interfaces ethernet eth0 address 10.0.0.1"
	var queue QueueItf
	queue = New()
	queue = queue.GetPathQue(path)
	queue.Print()
	fmt.Println(queue.GetPathString())
	fmt.Println(queue.GetFilePath())
}

func TestQueueNode_GetTop(t *testing.T) {
	path := "interface ethernet eth0 address"
	var queue QueueItf
	queue = New()
	queue.GetPathQue(path)
	queue.Print()
	fmt.Println(queue.GetPathString())
	//fmt.Println(queue.Peek())
}

func TestQueueNode_IsEmpty(t *testing.T) {
	path := "set interfaces 123 123 123"
	var queue QueueItf
	queue = new(Queue)
	queue = queue.GetPathQue(path)
	queue.Print()

	queue.Pop()
	queue.Print()
	fmt.Println(queue.IsEmpty())

	queue.Pop()
	queue.Print()
	fmt.Println(queue.IsEmpty())

	queue.Pop()
	queue.Pop()
	queue.Pop()

	fmt.Println(queue.IsEmpty())

}

func TestQueueNode_GetFilePath(t *testing.T) {
	path := "interfaces ethernet eth0 ip"
	var queue QueueItf
	queue = New()
	queue.GetPathQue(path)
	queue.Print()
	fmt.Println(queue.GetFilePath())
}

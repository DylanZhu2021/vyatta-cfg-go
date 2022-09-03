package gqueue

import (
	"container/list"
	"fmt"
	"github.com/gogf/gf/v2/container/glist"
	"vyatta-cfg-go/internal/service/logic/tools"
)

var _ QueueItf = new(Queue)

type Queue struct {
	q *glist.List
}

type QueueItf interface {
	Pop() interface{}
	Push(interface{})
	Length() int
	Print()
	Peek() interface{}
	IsEmpty() bool
	Iterator(f func(e *list.Element) bool)
	PathIntoQueue([]string)
	GetPathQue(string) *Queue
	GetPathString() string
	GetFilePath() string
}

func New() QueueItf {
	return &Queue{q: glist.New()}
}

// Length 链列个数
func (queue *Queue) Length() int {
	return queue.q.Len()
}

// Push 入列(insert)
func (queue *Queue) Push(data interface{}) {
	queue.q.PushBack(data)
}

// Pop 出队(delete)
func (queue *Queue) Pop() interface{} {
	return queue.q.PopFront()

}

func (queue *Queue) Peek() interface{} {
	return queue.q.Front().Value
}

// Print 打印链列
func (queue *Queue) Print() {

	queue.q.Iterator(func(e *glist.Element) bool {
		fmt.Println(e.Value)
		return true
	})

}

// Iterator 迭代器
func (queue *Queue) Iterator(f func(e *list.Element) bool) {
	queue.q.Iterator(f)
}

// IsEmpty 判断队空！！
func (queue *Queue) IsEmpty() bool {
	if queue.Length() == 0 {
		return true
	}
	return false
}

// PathIntoQueue 将path（[]string）转为queue
func (queue *Queue) PathIntoQueue(path []string) {

	//if len(path) > 1 {
	//	for i := 0; i < len(path); i++ {
	//		queue.Push(path[i])
	//	}
	//}
	for i := 0; i < len(path); i++ {
		queue.Push(path[i])
	}
}

// GetPathQue 把path由string转为queue
func (queue *Queue) GetPathQue(path string) *Queue {
	queue.PathIntoQueue(tools.GetPathArr(path))
	return queue
}

// GetPathString 把path（queue）转为string
func (queue *Queue) GetPathString() string {
	var path string
	path = queue.q.Join(" ")
	return path
}

// GetFilePath 把path(queue)转为文件完整的路径
//将tag节点路径转为含有node.tag的路径
//queue eg："interfaces ethernet eth0 address"
func (queue *Queue) GetFilePath() string {
	var path string
	var IsTag bool

	count := 0
	//使用queue的迭代器
	queue.Iterator(func(e *glist.Element) bool {
		if count == 0 {
			path += fmt.Sprintf("%v", e.Value)
			IsTag = tools.IsTag(path)
			count++
			return true
		}

		if !IsTag {
			path = path + "/" + fmt.Sprintf("%v", e.Value)
		} else {
			path += "/" + tools.NodeTag
		}
		IsTag = tools.IsTag(path)
		return true
	})
	return tools.TmpRootPath + "/" + path
}

//// NextIsEmpty 判断出一个之后是否为空队空！！
//func (queue *Queue) NextIsEmpty() bool {
//	queue.Pop()
//	if queue.Length() == 0 {
//		return true
//	}
//	return false
//}

//func (queue *Queue) GetSecond() string {
//	queue.Pop()
//	return fmt.Sprint(queue.Pop())
//}

//func (queue *Queue) GetThird() string {
//	return fmt.Sprint(queue.Next.Next.Next.Data)
//}

//// PopRetQue 出队(delete),返回剩下的队
//func (queue *Queue) PopRetQue() *Queue {
//	if queue == nil {
//		return nil
//	}
//	queue.Next = queue.Next.Next
//	return queue.Next
//
//}

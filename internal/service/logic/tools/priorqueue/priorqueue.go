package priorqueue

import (
	"container/heap"
	"container/list"
	"fmt"
	"vyatta-cfg-go/internal/service/logic/command/operatetree"
	"vyatta-cfg-go/internal/service/logic/tools"
	"vyatta-cfg-go/internal/service/logic/tools/gqueue"
)

type Item struct {
	Node  *operatetree.Node
	Index int
}

// PriorityQueue 优先级队列
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Node.Priority < pq[j].Node.Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Peek() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	return item
}

//修改优先级
func (pq *PriorityQueue) update(item *Item, priority int) {
	item.Node.Priority = priority
	heap.Fix(pq, item.Index)
}

func Print() {

	nodes := []*operatetree.Node{
		{
			"a",
			nil,
			nil,
			"",
			2,
			false,
			nil,
		},
		{
			"b",
			nil,
			nil,
			"",
			2,
			false,
			nil,
		},
		{
			"c",
			nil,
			nil,
			"",
			3,
			false,
			nil,
		},
		{
			"d",
			nil,
			nil,
			"",
			1,
			false,
			nil,
		},
	}

	pq := make(PriorityQueue, len(nodes))

	for i, node := range nodes {
		pq[i] = &Item{
			node,
			i,
		}

	}
	heap.Init(&pq)
	item := &Item{
		&operatetree.Node{
			"e",
			nil,
			nil,
			"",
			10,
			false,
			nil,
		},
		4,
	}

	heap.Push(&pq, item)

	pq.update(item, -1)
	for {
		if pq.Len() == 0 {
			break
		}
		item = heap.Pop(&pq).(*Item)
		fmt.Printf("index:%v value:%v  prio:%v\n", item.Index,
			item.Node.Name,
			item.Node.Priority)
	}

}

// CheckItem 检查Item中的node节点里面的set和del操作是不是符合要求
// 返回true为符合要求
// eg:set : 123,1234
//	  del : 123,234
// 此时对当前这个节点的set和del操作无效
// 实现逻辑：使用两个队列，将del的元素拿到set里卖弄去匹配，如果匹配不上则是无效
// 无值节点：
func (item Item) CheckItem() bool {
	var (
		setQue gqueue.QueueItf
		delQue gqueue.QueueItf
	)
	setQue = gqueue.New()
	delQue = gqueue.New()

	//如果Value不是空，则说明该节点不是无值节点
	if item.Node.Value != nil {
		for k, o := range item.Node.Operate {
			if o == "set" {
				setQue.Push(item.Node.Value[k])
			} else {
				delQue.Push(item.Node.Value[k])
			}
		}

		ret := true
		delQue.Iterator(func(delE *list.Element) bool {
			isExist := false
			for !setQue.IsEmpty() {
				if delE.Value == setQue.Pop() {
					isExist = true
					break
				}
			}
			//if not exist
			if !isExist {
				ret = true
				return true
			}
			item.Node.Value = tools.DeleteStrSliceElms(item.Node.Value, delE.Value.(string))
			return true
		})
		return ret
	} else {
		//如果为无值节点
		if item.Node.Operate != nil {
			for _, o := range item.Node.Operate {
				if o == "set" {
					setQue.Push(o)
				} else {
					delQue.Push(o)
				}
			}
			if delQue.Length() > setQue.Length() {
				item.Node.Operate = []string{
					"delete",
				}
				return true
			}
			if delQue.Length() == setQue.Length() {
				item.Node.Operate = []string{
					"set",
					"delete",
				}
				return true
			}
			if delQue.Length() < setQue.Length() {
				item.Node.Operate = []string{
					"set",
				}
				return true
			}
		}
	}
	return true
}

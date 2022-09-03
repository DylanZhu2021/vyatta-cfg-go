package commit

import (
	"fmt"
	"vyatta-cfg-go/internal/service/logic/activecfg"
	"vyatta-cfg-go/internal/service/logic/tmplate"
	"vyatta-cfg-go/internal/service/logic/tools/gqueue"
)

func Set(tree *activecfg.Node, path string, value string) error {

	var (
		isExist bool
		queue   gqueue.QueueItf
	)
	//初始化路劲队列
	queue = gqueue.New()
	queue.GetPathQue(path)

	tmpNode := tree

	for !queue.IsEmpty() {
		isExist = false
		//fmt.Println(queue.Peek())
		for _, node := range tmpNode.ChildNode {
			if node.Name == queue.Peek() {
				isExist = true
				tmpNode = node
				//queue.Pop()
				continue
			}

		}
		if !isExist {
			childnode := &activecfg.Node{
				Name: fmt.Sprintf("%v", queue.Peek()),
			}
			tmpNode.ChildNode = append(tmpNode.ChildNode, childnode)
			tmpNode = childnode
		}
		queue.Pop()

	}

	if value != "" {
		if tmplate.IsMulti(path) {
			tmpNode.Value = append(tmpNode.Value, value)
		} else {
			var valueArr []string
			valueArr = append(valueArr, value)
			tmpNode.Value = valueArr
		}

	}
	return nil
}

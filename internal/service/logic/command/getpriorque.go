package command

import (
	"container/heap"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"vyatta-cfg-go/internal/service/logic/command/operatetree"
	"vyatta-cfg-go/internal/service/logic/tools"
	"vyatta-cfg-go/internal/service/logic/tools/priorqueue"
)

// GetPriorQue (把tag节点和叶子节点放入优先级队列)
// 在过滤掉tag节点的时候，使用节点的完整路径判断
// 在过滤掉叶子节点的时候,通过diff树的形状判断！
func GetPriorQue(opTree *operatetree.Node) (priorqueue.PriorityQueue, error) {

	//由于优先级队列无法先指定大小，故先将过滤出来的节点放入[]*operatetree.Node中
	var priorNode *[]*operatetree.Node

	priorNode = new([]*operatetree.Node)

	//过滤
	GetPriorNode(opTree, priorNode)
	GetTagNode(opTree, priorNode)
	GetLeafNode(opTree, priorNode)

	//将结构体数组去重
	*priorNode = NodeArrDeduplicate(*priorNode)

	//初始化优先级队列
	priorQue := make(priorqueue.PriorityQueue, len(*priorNode))

	//节点入队
	i := 0
	for k, node := range *priorNode {
		item := &priorqueue.Item{
			node,
			i,
		}
		//如果是无效节点
		if !item.CheckItem() {
			return nil, errors.New(fmt.Sprintf("operate failed. path is: %v", item.Node.Path))
		}

		priorQue[k] = item
		i++
	}
	heap.Init(&priorQue)

	return priorQue, nil
}

// GetTagNode 过滤出tag节点
func GetTagNode(opTree *operatetree.Node, priorNode *[]*operatetree.Node) {

	if opTree == nil {
		return
	}

	pathArr := strings.Split(opTree.Path, " ")
	path := strings.Join(pathArr, "/")

	if tools.IsTag(path) {
		if opTree.GetPriority() == 0 {
			opTree.Priority = 10000
		}
		*priorNode = append(*priorNode, opTree)
	}
	//递归
	for _, node := range opTree.ChildNode {
		GetTagNode(node, priorNode)
	}

}

// GetLeafNode 过滤出叶子节点
func GetLeafNode(opTree *operatetree.Node, priorNode *[]*operatetree.Node) {
	if opTree == nil {
		return
	}
	if opTree.IsLeaf {
		if opTree.GetPriority() == 0 {
			opTree.Priority = 10000
		}
		*priorNode = append(*priorNode, opTree)
	}
	for _, node := range opTree.ChildNode {
		GetLeafNode(node, priorNode)
	}
}

// GetPriorNode 过滤出有优先级的节点
func GetPriorNode(opTree *operatetree.Node, priorNode *[]*operatetree.Node) {
	if opTree == nil {
		return
	}
	if opTree.GetPriority() != 0 {
		*priorNode = append(*priorNode, opTree)
	}
	for _, node := range opTree.ChildNode {
		GetPriorNode(node, priorNode)
	}
}

// NodeArrDeduplicate 结构体数组去重
// 使用reflect.DeepEqual函数
func NodeArrDeduplicate(priorNode []*operatetree.Node) (ret []*operatetree.Node) {
	n := len(priorNode)
	for i := 0; i < n; i++ {
		state := false
		for j := i + 1; j < n; j++ {
			if j > 0 && reflect.DeepEqual(priorNode[i], priorNode[j]) {
				state = true
				break
			}
		}
		if !state {
			ret = append(ret, priorNode[i])
		}
	}
	return
}

package commit

import (
	"container/heap"
	"errors"
	"fmt"
	"vyatta-cfg-go/internal/service/logic/activecfg"
	"vyatta-cfg-go/internal/service/logic/command"
	"vyatta-cfg-go/internal/service/logic/command/operatetree"
	"vyatta-cfg-go/internal/service/logic/datavalidate"
	"vyatta-cfg-go/internal/service/logic/tools/priorqueue"
)

/*
commit操作：  操作对象是cmds
1.将cmds转化为操作树：command函数
2.操作树过滤出节点放入优先级队列(如果操作树中的节点有不符合要求的直接设置失败)
3.优先级队列的节点合并到active config


详细：
1.commit操作要对同一个节点同一个值先set后del的这种节点要舍弃！
2.操作与值怎么对应？？
	单值节点？
	*
	多值节点？


流程：
1.先获取到cmds对应的操作树
2.然后获得优先级队列
3.优先级对列的节点转化为doCommitNode: value operate path
4.根据operate分为setQue和delQue两个队列，然后两个队列出队根据操作修改ACfg
5.

*/

// DoCommitNode 定义操作ACfg的节点！根据operate分类，根据path添加！
type DoCommitNode struct {
	Value   string `json:"value"`
	Operate string `json:"operate"`
	Path    string `json:"path"`
}

// Commit
// cmds => 操作树
// 操作树 => 优先级队列
// 优先级队列 => doCommitNode
// 最终执行在aCfg上
func Commit(cmds []string, aCfgTree *activecfg.Node) error {

	var tree *operatetree.Node
	tree = new(operatetree.Node)

	//生成操作树
	for _, cmd := range cmds {
		err := command.Command(tree, cmd)
		if err != nil {
			return err
		}
	}

	//生成优先级队列
	priorQue, err := command.GetPriorQue(tree)
	if err != nil {
		return err
	}

	for len(priorQue) > 0 {

		node := heap.Pop(&priorQue)

		//由优先级队列中的节点获取到doCommitNodes（执行commit操作的节点）
		doCommitNodes, err := PriorQueNodeToDoCommitNodes(node.(*priorqueue.Item))
		if doCommitNodes == nil && err == nil {
			continue
		}
		//将doCommitNode节点的操作整合到aCfg
		for _, doCommitNode := range doCommitNodes {
			ok, err := datavalidate.ValidateDate(doCommitNode.Path, doCommitNode.Value)
			if !ok {
				return errors.New(fmt.Sprintf("data:%s is not validate,error message:%s,path:%s",
					doCommitNode.Value, err, doCommitNode.Path))
			}
			err = ModifyActiveCfg(doCommitNode, aCfgTree)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// PriorQueNodeToDoCommitNodes 优先级队列中的一个节点转为一个或者多个(多值节点则一个值一个节点)DoCommitNode
func PriorQueNodeToDoCommitNodes(item *priorqueue.Item) ([]*DoCommitNode, error) {

	var doCommitNodes []*DoCommitNode

	if item.Node.Value == nil {
		if len(item.Node.Operate) == 1 {
			doCommitNodes = append(doCommitNodes, &DoCommitNode{
				"",
				item.Node.Operate[0],
				item.Node.Path,
			})
		} else {
			return nil, nil
		}
	}
	for k, v := range item.Node.Value {
		doCommitNodes = append(doCommitNodes, &DoCommitNode{
			v,
			item.Node.Operate[k],
			item.Node.Path,
		})
	}

	return doCommitNodes, nil
}

func ModifyActiveCfg(doCommitNode *DoCommitNode, aCfgTree *activecfg.Node) error {
	if doCommitNode.Operate == "set" {
		err := Set(aCfgTree, doCommitNode.Path, doCommitNode.Value)
		if err != nil {
			return err
		}
	} else {
		err := Delete(aCfgTree, doCommitNode.Path, doCommitNode.Value)
		if err != nil {
			return err
		}
	}
	return nil
}

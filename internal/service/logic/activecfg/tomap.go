package activecfg

// TreeToMap 将active config( struct 树型) => active config (map 树型)
func TreeToMap(tree *Node) (map[string]interface{}, []string) {

	//如果到了叶子节点就返回 value([]string)
	if tree.ChildNode == nil {
		return nil, tree.Value
	}

	var data map[string]interface{}
	data = make(map[string]interface{})

	//递归
	for _, node := range tree.ChildNode {
		tmpMap, valueArr := TreeToMap(node)
		if tmpMap != nil {
			data[node.Name] = tmpMap
		} else {
			data[node.Name] = valueArr
		}
	}

	return data, nil

}

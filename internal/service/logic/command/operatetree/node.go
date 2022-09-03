package operatetree

import "vyatta-cfg-go/internal/service/logic/tmplate"

type Node struct {
	Name      string   `json:"name"`
	Value     []string `json:"value"`
	Operate   []string `json:"operate"`
	Path      string   `json:"path"`
	Priority  int      `json:"priority"`
	IsLeaf    bool     `json:"isLeaf"`
	ChildNode []*Node  `json:"childNode"`
}

type OperateTree interface {
	GetPath() string
	GetPriority() int
	GetIsLeaf() bool
	GetOperateOfValue(int) string
	IsValueNode(string) (bool, error)
}

func (node Node) IsValueNode(path string) (bool, error) {
	return tmplate.IsNoValueNode(path)
}
func (node Node) GetPath() string {
	return node.Path
}

func (node Node) GetPriority() int {
	return node.Priority
}

func (node Node) GetIsLeaf() bool {
	return node.IsLeaf
}

// GetOperateOfValue 获取每个值对应的操作
// 比如set XXX 123
//    set XXX(和上面相同的路径) 1234  即将每个值和操作匹配起来
func (node Node) GetOperateOfValue(valueIndex int) string {
	return node.Operate[valueIndex]
}

//func GetJson() {
//	tree := &Node{
//		"interfaces",
//		nil,
//		[]*Node{
//			{
//				"ethernet",
//				[]string{
//					"eth0",
//				},
//
//				[]*Node{
//					{
//						"address",
//						[]string{
//							"dhcp",
//							"192.168.0.1/24",
//						},
//						nil,
//					},
//					{
//						"description",
//						[]string{
//							"OUTSIDE",
//						},
//						nil,
//					},
//					{
//						"hw-id",
//						[]string{
//							"00:0c:29:8f:74:b1",
//						},
//						nil,
//					},
//				},
//			},
//			{
//				"loopback",
//				[]string{
//					"lo",
//				},
//				nil,
//			},
//		},
//	}
//	jsonstr, _ := json.Marshal(tree)
//	fmt.Println(string(jsonstr))
//	var jsonnode *Node
//	err := json.Unmarshal(jsonstr, &jsonnode)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Printf("%v", jsonnode)
//}

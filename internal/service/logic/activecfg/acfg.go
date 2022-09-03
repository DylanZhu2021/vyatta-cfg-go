package activecfg

import "encoding/json"

// Node 定义active config node
type Node struct {
	Name      string   `json:"name"`
	Value     []string `json:"value"`
	ChildNode []*Node  `json:"childNode"`
}

func GetTree() string {
	//tree := &Node{
	//	"",
	//	nil,
	//	[]*Node{
	//		&Node{
	//			"interfaces",
	//			nil,
	//			[]*Node{
	//				&Node{
	//					"ethernet",
	//					nil,
	//					[]*Node{
	//						&Node{
	//							"eth0",
	//							nil,
	//							[]*Node{
	//								&Node{
	//									"address",
	//									[]string{
	//										"dhcp",
	//									},
	//									nil,
	//								},
	//							},
	//						},
	//						{
	//							"eth1",
	//							nil,
	//							[]*Node{
	//								{
	//									"address",
	//									[]string{
	//										"8.8.8.8/12",
	//									},
	//									nil,
	//								},
	//							},
	//						},
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	tree := &Node{}
	str, _ := json.Marshal(tree)
	return string(str)
}

package pattern

import (
	"fmt"
	"testing"
)

func TestCheckWithPattern(t *testing.T) {
	//完成四种测试！！！
	//traffic-policy shaper-hfsc node.tag class node.tag match
	path := "traffic-policy fair-queue" // /traffic-policy/random-detect/node.def:4:syntax:expression: pattern $VAR(@) "^[[:alnum:]][-_[:alnum:]]*$"

	//path := "traffic-policy round-robin node.tag class node.tag match" // /traffic-policy/round-robin/node.tag/class/node.tag/match/node.def:3:
	// syntax:expression: pattern $VAR(@) "^[^-]" ; "Match queue name cannot start with \"-\""
	//path := "cluster mcast-group" // /cluster/mcast-group/node.def:3:syntax:expression: pattern $VAR(@) "^239\."

	//path := "interfaces input" // interfaces/input/node.def:6:syntax:expression: pattern $VAR(@) "^ifb[0-9]+$" ; "name must be (ifb0-ifb999)"
	pattern, err := CheckWithPattern(path, "ifb123")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pattern)
}

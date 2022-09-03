package selectin

import (
	"fmt"
	"testing"
)

func TestGetSelect(t *testing.T) {
	path := "traffic-policy round-robin node.tag default queue-type"
	// traffic-policy/round-robin/node.tag/default/queue-type/node.def:3:
	//syntax:expression: $VAR(@) in "fq-codel", "fair-queue", "priority", "drop-tail";\
	selectin, _ := GetSelect(path)
	fmt.Println(selectin.Sel)
	fmt.Println(selectin.Ret)

}

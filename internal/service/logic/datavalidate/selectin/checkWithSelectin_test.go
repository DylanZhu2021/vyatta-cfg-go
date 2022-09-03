package selectin

import (
	"fmt"
	"testing"
)

func TestCheckWithSelectIn(t *testing.T) {
	//path := "traffic-policy round-robin node.tag default queue-type"
	// traffic-policy/round-robin/node.tag/default/queue-type/node.def:3:
	//syntax:expression: $VAR(@) in "fq-codel", "fair-queue", "priority", "drop-tail";\

	path := "load-balancing wan interface-health node.tag test node.tag type"
	// /load-balancing/wan/interface-health/node.tag/test/node.tag/type/node.def:5:
	// syntax:expression: $VAR(@) in "ping", "ttl", "user-defined";

	in, err := CheckWithSelectIn(path, "ttl")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(in)
}

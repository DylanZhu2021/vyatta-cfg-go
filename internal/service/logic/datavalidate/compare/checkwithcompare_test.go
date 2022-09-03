package compare

import (
	"fmt"
	"testing"
)

func TestCheckWithCompare(t *testing.T) {

	//完成测试！！！！
	//path := "load-balancing wan interface-health node.tag failure-count" // load-balancing/wan/interface-health/node.tag/failure-count/node.def:3:
	// syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 10; "failure count must be between 1-10"

	//path := "traffic-policy limiter node.tag class node.tag match node.tag ip max-length" // traffic-policy/limiter/node.tag/class/node.tag/match/node.tag/ip/max-length/node.def:3:
	// syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 65535; \

	//path := "system config-management commit-revisions" // system/config-management/commit-revisions/node.def:16:
	// syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 65535 ; \

	//path := "cluster monitor-dead-interval" // cluster/dead-interval/node.def:4:
	// syntax:expression: ($VAR(@) >= 1 && $VAR(@) < 30000) ; "Data must be between 1 and 29,999"

	//path := "traffic-policy fair-queue node.tag queue-limit" // /traffic-policy/fair-queue/node.tag/queue-limit/node.def:3:
	// syntax:expression: $VAR(@) > 1 && $VAR(@) < 128;\

	path := "traffic-policy random-detect node.tag precedence node.tag queue-limit" // traffic-policy/random-detect/node.tag/precedence/node.tag/queue-limit/node.def:3:
	// syntax:expression: $VAR(@) > 0 ; "Queue limit must greater than zero"

	compare, err := CheckWithCompare(path, "asd")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(compare)
}

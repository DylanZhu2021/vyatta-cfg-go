package tmplate

import (
	"fmt"
	"testing"
	"vyatta-cfg-go/internal/service/logic/tools/gqueue"
)

func TestGetTmpNode(t *testing.T) {
	//path := "interfaces loopback node.tag redirect" // /opt/vyatta/share/vyatta-cfg/templates/high-availability/virtual-server/node.tag/port/
	//path := "interfaces ethernet"
	//path := "high-availability vrrp group node.tag excluded-address"
	//path := "load-balancing wan interface-health node.tag success-count"
	path := "traffic-policy limiter node.tag class node.tag match node.tag ip max-length"
	//
	//path := "interfaces loopback node.tag redirect" // //opt/vyatta/share/vyatta-cfg/templates/interfaces/loopback/node.tag/redirect/
	//path := "traffic-policy network-emulator node.tag packet-reordering"
	var queue gqueue.QueueItf
	queue = gqueue.New()
	queue = queue.GetPathQue(path)

	tmpNode, err := GetTmpNode(queue)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", tmpNode.Syntax)
		fmt.Printf("%v\n", tmpNode.Commit)
		fmt.Printf("%v\n", tmpNode.Val_help)
	}
}

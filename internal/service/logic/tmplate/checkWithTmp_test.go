package tmplate

import (
	"fmt"
	"testing"
	"vyatta-cfg-go/internal/service/logic/tools/gqueue"
)

func TestCheckLeaf(t *testing.T) {
	path := "interfaces loopback node.tag redirect" //interfaces/loopback/node.tag/redirect/

	var queue gqueue.QueueItf
	queue = gqueue.New()
	queue.GetPathQue(path)
	queue.Print()

	fmt.Println(IsLeaf(queue.GetPathString()))
	fmt.Println(queue.GetFilePath())
}

package pattern

import (
	"fmt"
	"testing"
)

func TestGetPattern(t *testing.T) {
	path := "interfaces ethernet" // traffic-policy shaper-hfsc node.tag class node.tag match
	pattern, _ := GetPattern(path)
	fmt.Println(pattern.RegExp)
	fmt.Println(pattern.Ret)
}

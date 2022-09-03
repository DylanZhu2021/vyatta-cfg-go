package pl

import (
	"fmt"
	"testing"
)

func TestCheckWithPlScript(t *testing.T) {
	path := "traffic-policy shaper-hfsc node.tag bandwidth"
	script, err := CheckWithPlScript(path, "123")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(script)
}

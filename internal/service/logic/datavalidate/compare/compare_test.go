package compare

import (
	"fmt"
	"testing"
)

func TestGetCompare(t *testing.T) {
	path := "cluster monitor-dead-interval"
	compare, err := GetCompare(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(compare.Ret)
}

package get

import (
	"fmt"
	"testing"
)

func TestGetSyntaxExp(t *testing.T) {

	path := "cluster dead-interval" // /opt/vyatta/share/vyatta-cfg/templates/cluster/dead-interval/
	syntax, err := GetSyntaxExp(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(syntax)
}

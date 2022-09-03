package script

import (
	"fmt"
	"testing"
)

func TestGetValidate(t *testing.T) {
	path := "interfaces pseudo-ethernet node.tag mode/" // /opt/vyatta/share/vyatta-cfg/templates/interfaces/pseudo-ethernet/node.tag/mode/
	valid, err := GetValidate(path, "private")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(valid.Exec)
	fmt.Println(valid.Ret)

}

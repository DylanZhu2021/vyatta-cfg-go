package command

import (
	"fmt"
	"testing"
)

func TestGetCmdValue(t *testing.T) {
	setCmd := "set system host-name 123zxzx"
	fmt.Println(GetCmdValue(setCmd, false))
}

func TestGetPriority(t *testing.T) {
	prio := GetPriority("interfaces ethernet")
	fmt.Println(prio)
}

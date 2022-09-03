package script

import (
	"fmt"
	"testing"
)

func TestCheckWithScript(t *testing.T) {

	//完成测试！！！！
	// firewall/group/ipv6-address-group/node.tag/address/node.def:6:
	//syntax:expression: exec "${vyos_libexec_dir}/validate-value  --exec \"${vyos_validators_dir}/ipv6-address \" --exec \"${vyos_validators_dir}/ipv 6-range \"  --value \'$VAR(@)\'"; "Invalid value"

	path := "firewall group ipv6-address-group node.tag address"
	//path := "policy access-list6 node.tag rule node.tag action" // firewall broadcast-ping
	script, err := CheckWithScript(path, "123") //enable
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(script)
}

package datavalidate

import (
	"fmt"
	"testing"
)

func TestExecCommand(t *testing.T) {
	//strData := ExecCommand("/usr/libexec/vyos/validate-value --regex \\'(dhcp|dhcpv6)\\' --exec \\\"/usr/libexec/vyos/ip-host \\\"  --value \\'123\\'")
	strData := ExecCommand("${vyos_libexec_dir}/validate-value \n\n --exec \\\"${vyos_validators_dir}/numeric\n\n --range 1-4294967294\\\"  \n\n--value \\'$VAR(@)\\'")
	//${vyos_libexec_dir}/validate-value  --exec \"${vyos_validators_dir}/numeric --range 1-999999\"  --value \'$VAR(@)\'
	//strData := "${vyos_libexec_dir}/validate-value  --exec \"${vyos_validators_dir}/numeric --range 1-999999\"  --value "
	fmt.Println("execute finished:" + strData)
}

func TestReplace(t *testing.T) {
	Replace()
}

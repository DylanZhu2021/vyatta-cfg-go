package pl

import (
	"errors"
	"os/exec"
	"regexp"
	"strings"
	"vyatta-cfg-go/internal/service/logic/datavalidate/get"
)

/*
/opt/vyatta/share/vyatta-cfg/templates/traffic-policy/shaper-hfsc/node.tag/bandwidth/
syntax:expression: $VAR(@) == "auto" || \
		   exec "/opt/vyatta/sbin/vyatta-qos-util.pl --rate $VAR(@)"
*/

func CheckWithPlScript(path string, value string) (bool, error) {

	if value == "auto" {
		return true, nil
	}
	syntax, err := get.GetSyntaxExp(path)
	if err != nil {
		return false, err
	}

	//获取到pl脚本执行语句
	execStr := strings.Split(strings.Split(syntax, "||")[1], "\"")[1]

	//替换掉$VAR(@)为需要设置的值
	reg := regexp.MustCompile(`\$VAR\(\@\)`)
	execStr = reg.ReplaceAllString(execStr, value)

	cmd := exec.Command("/bin/bash", "-c", "-lc", execStr) // /home/vyos/vyatta-cfg-go  "/home/vyos/vyatta-cfg-go/env.sh"
	out, _ := cmd.CombinedOutput()

	if string(out) == "" {
		return true, nil
	}
	return false, errors.New(string(out))
}

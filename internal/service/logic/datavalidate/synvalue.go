package datavalidate

import (
	"fmt"
	"os/exec"
	"regexp"
	"vyatta-cfg-go/internal/service/logic/datavalidate/get"
	datavalidate "vyatta-cfg-go/internal/service/logic/datavalidate/script"
)

// ExecCommand go执行命令，并返回值
func ExecCommand(strCommand string) string {
	cmd := exec.Command("/bin/bash", "-c", strCommand)

	out, _ := cmd.CombinedOutput()
	return string(out)
}

//exec "${vyos_libexec_dir}/validate-value
//--exec \"${vyos_validators_dir}/numeric
//--range 1-4294967294\"
//--value \'$VAR(@)\'"; "Invalid value"

// Replace 正则匹配与替换
func Replace() {
	str := "exec \"${vyos_libexec_dir}/validate-value  " +
		"--exec \\\"${vyos_validators_dir}/numeric " +
		"--range 1-4294967294\\\"  --value \\'$VAR(@)\\'\"; \"Invalid value\"\n"

	//reg := regexp.MustCompile(`\$VAR\(\@\)`)
	//str = reg.ReplaceAllString(str, "123")
	//fmt.Println(str)
	//
	reg := regexp.MustCompile(`\$\{\w*\}`)

	arr := reg.FindAll([]byte(str), -1)
	fmt.Println(string(arr[0]))
	fmt.Println(get.GetEnvName(string(arr[0])))
	fmt.Println(datavalidate.GetEnvValue(get.GetEnvName(string(arr[0]))))
	fmt.Println(string(arr[1]))
	fmt.Println(get.GetEnvName(string(arr[1])))

	envValue, err := datavalidate.GetEnvValue(get.GetEnvName(string(arr[0])))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reg.ReplaceAllString(str, envValue))
}

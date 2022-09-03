package script

import (
	"regexp"
	"strings"
	"vyatta-cfg-go/internal/service/logic/datavalidate/get"
)

/*
这个是校验的文件
*/

type Script struct {
	Exec string
	Ret  string
}

// GetValidate 根据expression获取validate结构体
/*
exec "${vyos_libexec_dir}/validate-value
						--exec \"${vyos_validators_dir}/numeric --range 1-4294967294\"
						--value \'$VAR(@)\'"; "Invalid value"
*/
func GetValidate(path string, value string) (*Script, error) {
	var script *Script

	script = new(Script)
	//获取syntax
	syntax, err := get.GetSyntaxExp(path)
	if err != nil {
		return nil, err
	}

	strSplit := strings.Split(syntax, ";")
	//获取到Ret
	retArr := strings.Split(strSplit[1], "\"")

	script.Ret = retArr[1]

	//获取Exec
	execArr := strings.Split(strSplit[0], "\"")

	exec := ""
	for i, v := range execArr {
		if i == 0 {
			continue
		}
		if exec == "" {
			exec += v
			continue
		}
		if v == "" {
			continue
		}
		exec += "\"" + v
	}

	//替换环境变量
	execRepEnv, err := RegReplEnv(exec)
	if err != nil {
		return nil, err
	}

	//替换值
	execRepVal := RegReplValue(execRepEnv, value)

	//删除字符串execRepVal中的反斜杠
	strArr := strings.Split(execRepVal, "\\")
	execRepVal = strings.Join(strArr, "")

	script.Exec = execRepVal

	return script, nil
}

// RegReplEnv 将exec中的环境遍历替换掉
func RegReplEnv(exec string) (string, error) {

	var (
		err      error
		envValue []string
	)

	//正则匹配，获取到环境变量的名字
	reg := regexp.MustCompile(`\$\{\w*\}`)

	arr := reg.FindAllString(exec, -1)
	envValue = make([]string, len(arr))

	//获取到exec所需要的环境变量值
	i := 0
	for _, v := range arr {
		envValue[i], err = GetEnvValue(get.GetEnvName(v))
		if err != nil {
			return "", err
		}
		i++
	}

	//将exec里面的环境变量替换为值
	for k, env := range envValue {
		exec = strings.Replace(exec, arr[k], env, -1)
	}

	return exec, nil

}

// RegReplValue 将exec中的值替换掉
func RegReplValue(exec string, value string) string {

	reg := regexp.MustCompile(`\$VAR\(\@\)`)

	return reg.ReplaceAllString(exec, value)
}

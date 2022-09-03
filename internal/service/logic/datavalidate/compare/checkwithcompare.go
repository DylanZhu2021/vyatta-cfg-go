package compare

import (
	"errors"
	"github.com/PaulXu-cn/goeval"
	"regexp"
)

func CheckWithCompare(path string, value string) (bool, error) {

	//$VAR(@) >= 1 && $VAR(@) < 30000
	//获取到compare结构体
	compare, err := GetCompare(path)
	if err != nil {
		return false, err
	}

	//正则替换掉$VAR(@),换为需要校验的值！
	reg := regexp.MustCompile(`\$VAR\(\@\)`)
	compare.ComExp = reg.ReplaceAllString(compare.ComExp, value)

	execStr := "fmt.Print(" + compare.ComExp + ")"

	//将字符串当成go语句执行，然后返回结果！！
	if re, err := goeval.Eval("", execStr, "fmt"); nil == err {
		if string(re) == "true" {
			return true, nil
		}
		return false, errors.New(compare.Ret)
	} else {
		return false, err
	}
}

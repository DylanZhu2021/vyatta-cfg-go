package pattern

import (
	"errors"
	"regexp"
)

func CheckWithPattern(path string, value string) (bool, error) {

	//根据路径，path，获取到pattern（根据syntax字段转换成的结构体）
	pattern, err := GetPattern(path)
	if err != nil {
		return false, err
	}

	//将正则表达式与需要设置的值进行匹配
	match, err := regexp.MatchString(pattern.RegExp, value)
	if err != nil {
		return false, err
	}

	if !match {
		return false, errors.New(pattern.Ret)
	}

	return match, nil
}

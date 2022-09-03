package datavalidate

import (
	"errors"
	"fmt"
	"regexp"
	"vyatta-cfg-go/internal/service/logic/datavalidate/compare"
	"vyatta-cfg-go/internal/service/logic/datavalidate/get"
	"vyatta-cfg-go/internal/service/logic/datavalidate/pattern"
	"vyatta-cfg-go/internal/service/logic/datavalidate/pl"
	"vyatta-cfg-go/internal/service/logic/datavalidate/script"
	"vyatta-cfg-go/internal/service/logic/datavalidate/selectin"
)

// ValidateDate 数据校验
// 返回true，表示数据合格
func ValidateDate(path string, value string) (bool, error) {

	syntax, err := get.GetSyntaxExp(path)
	if err != nil {
		return false, errors.New(fmt.Sprintf("get syntax fail,err:%v", err))
	}

	//匹配pl
	reg, err := regexp.Compile(".auto.")
	if err != nil {
		return false, err
	}
	if reg.MatchString(syntax) {
		fmt.Println("校验规则：pl脚本")
		ret, err := pl.CheckWithPlScript(path, value)
		if err != nil {
			return false, err
		}
		return ret, nil
	}

	//匹配exec执行脚本
	reg, err = regexp.Compile("exec.")
	if err != nil {
		return false, err
	}
	if reg.MatchString(syntax) {
		fmt.Println("校验规则：exec执行脚本")
		ret, err := script.CheckWithScript(path, value)
		if err != nil {
			return false, err
		}
		return ret, nil
	}

	//匹配pattern模式
	reg, err = regexp.Compile("pattern.")
	if err != nil {
		return false, err
	}
	if reg.MatchString(syntax) {
		fmt.Println("校验规则：pattern模式")
		ret, err := pattern.CheckWithPattern(path, value)
		if err != nil {
			return false, err
		}
		return ret, nil
	}

	//匹配compare
	reg, err = regexp.Compile(".\\$VAR\\(\\@\\)[\\ ][>|<|>=|<=].")
	if err != nil {
		return false, err
	}
	if reg.MatchString(syntax) {
		fmt.Println("校验规则：compare")
		ret, err := compare.CheckWithCompare(path, value)
		if err != nil {
			return false, err
		}
		return ret, nil
	}

	//匹配selectin
	reg, err = regexp.Compile("\\$VAR\\(\\@\\)\\ in.")
	if err != nil {
		return false, err
	}
	if reg.MatchString(syntax) {
		fmt.Println("校验规则：selectin")
		ret, err := selectin.CheckWithSelectIn(path, value)
		if err != nil {
			return false, err
		}
		return ret, nil
	}

	return false, errors.New("not find check mode")
}

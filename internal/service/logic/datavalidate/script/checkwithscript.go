package script

import (
	"errors"
	"os/exec"
)

// CheckWithScript 提供与脚本进行校验的函数接口
func CheckWithScript(path string, value string) (bool, error) {

	//获取到validate结构体
	valid, err := GetValidate(path, value)
	if err != nil {
		return false, err
	}

	cmd := exec.Command("/bin/bash", "-c", valid.Exec)

	data, _ := cmd.CombinedOutput()

	if string(data) == "" {
		return true, nil
	}

	return false, errors.New(valid.Ret)
}

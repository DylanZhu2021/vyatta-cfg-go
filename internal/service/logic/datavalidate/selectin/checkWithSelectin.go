package selectin

import "errors"

func CheckWithSelectIn(path string, value string) (bool, error) {

	selectin, err := GetSelect(path)
	if err != nil {
		return false, err
	}
	for i := 0; i < len(selectin.Sel); i++ {
		if selectin.Sel[i] == value {
			return true, nil
		}
	}

	return false, errors.New(selectin.Ret)
}

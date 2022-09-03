package script

import (
	"fmt"
	"testing"
)

func TestGetEnvValueFromFile(t *testing.T) {
	emap := GetEnvValueFromFile()
	fmt.Println(len(emap))
}

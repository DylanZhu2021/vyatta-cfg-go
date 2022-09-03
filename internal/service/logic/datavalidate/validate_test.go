package datavalidate

import (
	"fmt"
	"testing"
)

func TestValidDate(t *testing.T) {
	//path := "firewall group ipv6-address-group node.tag address"
	//path := "load-balancing wan interface-health node.tag test node.tag type"
	//path := "traffic-policy shaper-hfsc node.tag bandwidth"
	//path := "traffic-policy fair-queue"
	//path := "traffic-policy random-detect node.tag precedenc node.tag queue-limit"
	path := "system name-server"
	date, err := ValidateDate(path, "0.0.0.0/12")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(date)
}

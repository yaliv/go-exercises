package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	
	// return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
	
	var dottedIP string
	// Assume we don't know the number of IP numsets.
	for i, numset := range ip {
		if i > 0 { dottedIP += "." }
		dottedIP += strconv.Itoa(int(numset))
	}
	return fmt.Sprintf("%q", dottedIP)
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

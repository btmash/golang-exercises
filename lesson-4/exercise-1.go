package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	// Make might be preferable for arrays, maps.
	// ipString := make([]string, len(ip))
	ipString := [len(ip)]string{}
	for i, v := range ip {
		ipString[i] = strconv.Itoa(int(v))
	}
	ipSlice := ipString[:]
	// string.Join works on slices, not arrays.
	return strings.Join(ipSlice, ".")
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

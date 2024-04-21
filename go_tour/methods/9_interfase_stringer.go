package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	arr2 := make([]string, len(ip))

	for idx, i := range [len(ip)]byte(ip) {
		// arr2 = append(arr2, fmt.Sprint(i))
		// arr2[idx] = strconv.Itoa(int(i))
		arr2[idx] = fmt.Sprint(i)
	}
	str_arr := strings.Join(arr2, ".")
	return str_arr
}

// 	// s := make([]string, len(ip))

// 	// for i, val := range ip {
// 	// s[i] = strconv.Itoa(int(val))
// 	// 	s[i] = val
// 	// }
// 	// return strings.Join(s, ".")
// 	return strings.Join(strings.Split((strings.Trim(fmt.Sprint([len(ip)]byte(ip)), "[]")), " "), ".")
// }

// TODO: Add a "String() string" method to IPAddr.

func main() {
	// hosts := map[string]IPAddr{
	// 	"loopback":  {127, 0, 0, 1},
	// 	"googleDNS": {8, 8, 8, 8},
	// }
	// for name, ip := range hosts {
	// 	fmt.Printf("%v: %v\n", name, ip)
	// }
	ip := IPAddr{127, 0, 1, 1}
	fmt.Print(string(ip))
	// fmt.Println(ip)

}

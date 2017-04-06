package main

import "fmt"
import "net"

import "gopkg.in/alecthomas/kingpin.v2"

var (
	cidr = kingpin.Arg("cidr", "Specify CIDR for range").Required().String()
)

func main() {
	kingpin.Parse()

	ip, ipnet, _ := net.ParseCIDR(*cidr)
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		fmt.Println(ip)
	}
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

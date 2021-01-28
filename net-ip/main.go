package main

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/signalsciences/ipv4"
)

var privateIPBlocks []*net.IPNet

func init() {
	for _, cidr := range []string{
		"127.0.0.0/8",    // IPv4 loopback
		"10.0.0.0/8",     // RFC1918
		"172.16.0.0/12",  // RFC1918
		"192.168.0.0/16", // RFC1918
		"169.254.0.0/16", // RFC3927 link-local
		"::1/128",        // IPv6 loopback
		"fe80::/10",      // IPv6 link-local
		"fc00::/7",       // IPv6 unique local addr
	} {
		_, block, err := net.ParseCIDR(cidr)
		if err != nil {
			panic(fmt.Errorf("parse error on %q: %v", cidr, err))
		}
		privateIPBlocks = append(privateIPBlocks, block)
	}
}

func isPrivateIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() {
		return true
	}

	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}

func ipToInt(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func main() {
	fmt.Println(isPrivateIP(net.ParseIP("172.28.4.60")))
	fmt.Println(ipToInt(net.ParseIP("2409:8a34:7101:330c:20bf:688f:285f:5b76")))
	fmt.Println(ipToInt(net.ParseIP("172.28.4.60")))

	ipNum, err := ipv4.FromDots("172.28.4.60")

	if err != nil {
		panic(err)
	}

	fmt.Println("ip v4 to int", ipNum)
}

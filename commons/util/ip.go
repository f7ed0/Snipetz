package util

import (
	"net"
)

func GetIPs() (adresses map[string][]net.IP, err error) {
	adresses = make(map[string][]net.IP)
	ifaces, err := net.Interfaces()
	if err != nil {
		return
	}
	// handle err
	for _, i := range ifaces {
		adresses[i.Name] = make([]net.IP, 0)
		var addrs []net.Addr
		addrs, err = i.Addrs()
		if err != nil {
			return
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			adresses[i.Name] = append(adresses[i.Name], ip)
		}
	}
	return

}

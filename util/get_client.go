package util

import (
	"fmt"
	"net"
)

// GetMac 获取mac地址
func GetMac() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v\n", err)
		return macAddrs
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		macAddrs = append(macAddrs, macAddr)
	}
	return macAddrs
}

// GetLocalMac 获取本地mac地址
func GetLocalMac() (mac string) {
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("Poor soul, here is what you got: " + err.Error())
	}
	for _, inter := range interfaces {
		fmt.Println(inter.Name)
		mac := inter.HardwareAddr //获取本机MAC地址
		fmt.Println("MAC = ", mac)
	}
	return mac
}

// GetIps 获取本机ip [192.168.1.11]
func GetIps() (ips []string) {
	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("fail to get net interfaces ipAddress: %v\n", err)
		return ips
	}

	for _, address := range interfaceAddr {
		ipNet, isVailIpNet := address.(*net.IPNet)
		// 检查ip地址判断是否回环地址
		if isVailIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}

// GetExternal 获取外网ip
//func GetExternal() {
//	resp, err := http.Get("https://myexternalip.com/raw")
//	if err != nil {
//		os.Stderr([]byte(err.Error() + "\n")) // Convert string to byte slice
//		os.Exit(1)
//	}
//	defer resp.Body.Close()
//
//	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
//		os.Stderr([]byte(err.Error() + "\n"))
//		os.Exit(1)
//	}
//	os.Exit(0)
//}

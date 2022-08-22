package util

import (
	"errors"
	"fmt"
	"net"
)

func GetLocalIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("get fail")
}

func GetIp(host string, port int) (string, error) {
	addr := fmt.Sprint(host, ":", port)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if nil != err {
		return "", err
	}
	return tcpAddr.IP.String(), nil
}

func GetIp4(host string, port string) (string, error) {
	addr := fmt.Sprint(host, ":", port)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if nil != err {
		return "", err
	}
	return tcpAddr.IP.String(), nil
}

// 获取空闲端口
func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

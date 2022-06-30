package extip

import (
	"errors"
	"github.com/jageros/hawox/httpc"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
)

func GetMyOutIp() (string, error) {
	resp, err := httpc.Request(httpc.GET, "http://idata.hawtech.cn/my-ip", httpc.FORM, nil, nil)
	if err != nil {
		return "", err
	}
	ip := string(resp)
	return ip, nil
}

// GetExternalIP from https://ifconfig.co/ip
func GetExternalIP() (string, error) {
	resp, err := http.Get("https://ifconfig.co/ip")
	if err != nil {
		log.Printf("http get err: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("err: %v", err)
	}
	ips := strings.Split(string(body), "\n")
	ipStr := ips[0]
	return ipStr, nil
}

// ====== from https://2020.ip138.com ======
func GetExtIP() (string, error) {
	resp, err := http.Get("https://2021.ip138.com")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var ipStr string
	for _, line := range strings.Split(string(body), "\n") {
		log.Println(line)
		if strings.HasPrefix(line, "<title>您的IP地址是") {
			ipStr = strings.Split(line[27:], "<")[0]
			break
		}
		if strings.HasPrefix(line, "您的iP地址是") {
			ipStr = strings.Split(line[21:], "]")[0]
			break
		}
	}
	return ipStr, nil
}

func GetMyIp(netCard string) (string, error) {
	ift, err := net.InterfaceByName(netCard)
	if err != nil {
		return "", err
	}
	addrs, err := ift.Addrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("NotExistIpErr")
}

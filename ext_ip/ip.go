package extip

import (
	"github.com/jageros/hawox/httpc"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	ipAddr string
)

func GetMyIp() (string, bool) {
	resp, err := httpc.Request(httpc.GET, "http://idata.hawtech.cn/my-ip", httpc.FORM, nil, nil)
	if err != nil {
		return ipAddr, false
	}
	ip := string(resp)
	if ipAddr != ip {
		ipAddr = ip
		return ip, true
	}
	return ip, false
}

// ====== from https://ifconfig.co/ip ======
func GetExternalIP() (string, bool) {
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
	if ipAddr != ipStr {
		log.Printf("Get ExternalIP=%s oldIP=%s", ipStr, ipAddr)
		ipAddr = ipStr
		return ipStr, true
	}
	return ipStr, false
}

// ====== from https://2020.ip138.com ======
func GetExtIP() (string, bool) {
	resp, err := http.Get("https://2021.ip138.com")
	if err != nil {
		log.Printf("http get err: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("err: %v", err)
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
	if ipAddr != ipStr {
		log.Printf("Get ExternalIP=%s oldIP=%s", ipStr, ipAddr)
		ipAddr = ipStr
		return ipStr, true
	}
	return ipStr, false
}

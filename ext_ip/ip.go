package extip

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var ipAddr string

// ====== from http://ipinfo.io/ip ======
func GetExternalIP() (string, bool) {
	resp, err := http.Get("http://ipinfo.io/ip")
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
	resp, err := http.Get("https://2020.ip138.com")
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

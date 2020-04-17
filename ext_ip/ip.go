package extip

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var ipAddr string

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
	log.Printf("Get ExternalIP=%s oldIP=%s", ipStr, ipAddr)
	if ipAddr != ipStr {
		ipAddr = ipStr
		return ipStr, true
	}
	return ipStr, false
}

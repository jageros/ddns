package main

import (
	"ddns_pro/consts"
	"ddns_pro/ddns"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	resp, err := http.Get("http://ipinfo.io/ip")
	if err != nil {
		log.Printf("http get err: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	fmt.Printf(string(body))
	ips := strings.Split(string(body), "\n")
	ipAddr := ips[0]

	// ===========
	recordId := "3j4h5tf8d9ij"
	key2val := map[string]string{
		"Action":          "RecordModify",
		"SecretId":        consts.SecretId,
		"Timestamp":       strconv.FormatInt(time.Now().Unix(), 10),
		"Nonce":           strconv.Itoa(rand.Intn(10000)),
		"SignatureMethod": "HmacSHA256",
		"domain":          "hawtech.cn",
		"recordId":        recordId,
		"subDomain":       "pi",
		"recordType":      "CNAME",
		"recordLine":      "默认",
		"value":           ipAddr,
	}
	ags := ddns.NewDDnsArgSt(key2val)
	ags.GenSignature()
	ags.UrlEncode()
	urlStr := ags.GenUrl()
	fmt.Println(urlStr)
}

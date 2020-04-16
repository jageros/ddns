package main

import (
	"ddns_pro/consts"
	"ddns_pro/ddns"
	"encoding/json"
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
	recordId := getRecordId()
	fmt.Println(recordId)
	key2val := map[string]string{
		"Action":          "RecordModify",
		"SecretId":        consts.SecretId,
		"Timestamp":       strconv.FormatInt(time.Now().Unix(), 10),
		"Nonce":           strconv.Itoa(rand.Intn(10000)),
		"SignatureMethod": "HmacSHA256",
		"domain":          "hawtech.cn",
		"recordId":        recordId,
		"subDomain":       "pi",
		"recordType":      "A",
		"recordLine":      "默认",
		"value":           ipAddr,
	}
	ags := ddns.NewDDnsArgSt(key2val)
	ags.GenSignature()
	ags.UrlEncode()
	urlStr := ags.GenUrl(false)
	fmt.Println(urlStr)

	resp, err = http.Get(urlStr)
	if err != nil {
		log.Printf("http get err: %v", err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read body err: %v", err)
	}
	fmt.Println(string(body))
}

func getRecordId() string {
	key2val := map[string]string{
		"Action":          "RecordList",
		"SecretId":        consts.SecretId,
		"Timestamp":       strconv.FormatInt(time.Now().Unix(), 10),
		"Nonce":           strconv.Itoa(rand.Intn(10000)),
		"SignatureMethod": "HmacSHA256",
		"domain":          "hawtech.cn",
		"subDomain":       "pi",
	}
	ags := ddns.NewDDnsArgSt(key2val)
	ags.GenSignature()
	ags.UrlEncode()
	urlStr := ags.GenUrl(false)
	fmt.Println(urlStr)
	resp, err := http.Get(urlStr)
	if err != nil {
		log.Printf("http get err: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read body err: %v", err)
	}
	var reply = replyData{}
	err = json.Unmarshal(body, &reply)
	if err != nil {
		log.Printf("json.Unmarshal err: %v", err)
	}
	return strconv.Itoa(reply.Data.Records[0].ID)
}

type replyData struct {
	Data data `json:"data"`
}

type data struct {
	Records []record `json:"records"`
}

type record struct {
	ID int `json:"id"`
}

package ddns

import (
	"ddns_pro/consts"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func getRecordId(subDomain string) string {
	key2val := map[string]string{
		"Action":          "RecordList",
		"SecretId":        consts.SecretId,
		"Timestamp":       strconv.FormatInt(time.Now().Unix(), 10),
		"Nonce":           strconv.Itoa(rand.Intn(10000)),
		"SignatureMethod": "HmacSHA256",
		"domain":          consts.Domain,
		"subDomain":       subDomain,
	}
	ags := newDDnsArgSt(key2val)
	ags.GenSignature()
	ags.UrlEncode()
	urlStr := ags.GenUrl(false)
	resp, err := http.Get(urlStr)
	if err != nil {
		log.Printf("http get err: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read body err: %v", err)
	}
	var reply = map[string]interface{}{}
	err = json.Unmarshal(body, &reply)
	if err != nil {
		log.Printf("json.Unmarshal err: %v", err)
	}
	return getID(reply)
}

func getID(arg map[string]interface{}) string {
	if arg["code"].(float64) != 0 {
		log.Printf("get record list err, reply=%v", arg)
	}
	records := arg["data"].(map[string]interface{})["records"].([]interface{})
	if len(records) <= 0 {
		return ""
	}
	arg1, ok := records[0].(map[string]interface{})["id"].(float64)
	if !ok {
		return ""
	}
	return strconv.Itoa(int(arg1))
}

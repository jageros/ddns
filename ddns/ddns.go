package ddns

import (
	"crypto/hmac"
	"crypto/sha256"
	"ddns_pro/consts"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

type arg struct {
	key   string
	value string
}

type dDnsArgSt struct {
	args []arg
}

// ================== sort interface ====================
func (d *dDnsArgSt) Len() int {
	return len(d.args)
}

func (d *dDnsArgSt) Less(i, j int) bool {
	return d.args[i].key < d.args[j].key
}

func (d *dDnsArgSt) Swap(i, j int) {
	d.args[i], d.args[j] = d.args[j], d.args[i]
}

// ========================================================

func newDDnsArgSt(ms map[string]string) *dDnsArgSt {
	da := &dDnsArgSt{}
	for key, val := range ms {
		da.args = append(da.args, arg{
			key:   key,
			value: val,
		})
	}
	sort.Sort(da)
	return da
}

func (d *dDnsArgSt) UrlEncode() {
	for i, a := range d.args {
		d.args[i].value = url.QueryEscape(a.value)
	}
}

func (d *dDnsArgSt) GenUrl(isSig bool) string {
	urlStr := consts.BaseUrl
	if isSig {
		urlStr = consts.SigBaseUrl
	}
	for i, a := range d.args {
		field := "&"
		if i <= 0 {
			field = "?"
		}
		urlStr += fmt.Sprintf("%s%s=%s", field, a.key, a.value)
	}
	return urlStr
}

func (d *dDnsArgSt) GenSignature() {
	urlStr := d.GenUrl(true)
	h := hmac.New(sha256.New, []byte(consts.SecretKey))
	_, err := h.Write([]byte(urlStr))
	if err != nil {
		log.Printf("write err: %v", err)
	}
	sig := base64.StdEncoding.EncodeToString(h.Sum(nil))
	d.args = append(d.args, arg{
		key:   "Signature",
		value: sig,
	})
}

// ================

func SetDns(ipAddr string) {
	key2val := map[string]string{
		"Action":          "RecordModify",
		"SecretId":        consts.SecretId,
		"Timestamp":       strconv.FormatInt(time.Now().Unix(), 10),
		"Nonce":           strconv.Itoa(rand.Intn(10000)),
		"SignatureMethod": "HmacSHA256",
		"domain":          consts.Domain,
		"recordId":        getRecordId(),
		"subDomain":       consts.SubDomain,
		"recordType":      "A",
		"recordLine":      "默认",
		"value":           ipAddr,
	}
	ags := newDDnsArgSt(key2val)
	ags.GenSignature()
	ags.UrlEncode()
	urlStr := ags.GenUrl(false)
	//fmt.Println(urlStr)

	resp, err := http.Get(urlStr)
	if err != nil {
		log.Printf("http get err: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read body err: %v", err)
	}
	reply := map[string]interface{}{}
	err = json.Unmarshal(body, &reply)
	code := reply["code"].(float64)
	log.Printf("Update Dns status code=%d", int(code))
}

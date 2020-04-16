package ddns

import (
	"crypto/hmac"
	"crypto/sha256"
	"ddns_pro/consts"
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"sort"
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

func NewDDnsArgSt(ms map[string]string) *dDnsArgSt {
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

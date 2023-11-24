package main

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

func main() {
	signer := NewSigner("pn5d2bef5e37774381")
	redirectURI, _ := url.Parse("")
	encodedURI := url.QueryEscape(redirectURI.String())
	params := map[string]string{"uid": "30", "timestamp": "1700794400", "ticket": "420c3630a3d468ab8dc1807729eed784", "redirect_uri": encodedURI, "nickname": "18983663382"}
	sign := signer.Sign(params)
	fmt.Println(sign)
}

type Signer struct{ appKey string }

func NewSigner(appKey string) *Signer { return &Signer{appKey: appKey} }

func (s *Signer) Sign(params map[string]string) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var paramsStr strings.Builder
	for _, k := range keys {
		paramsStr.WriteString(k)
		paramsStr.WriteString("=")
		paramsStr.WriteString(params[k])
		paramsStr.WriteString("&")
	}
	paramsStr.WriteString("app_key=")
	paramsStr.WriteString(s.appKey)
	fmt.Println("---", paramsStr.String())
	hash := md5.Sum([]byte(paramsStr.String()))
	return fmt.Sprintf("%x", hash)
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	var data = strings.NewReader(`{"Gjbq":"1","FpztDm":["01"],"FplyDm":"0","FplxDm":[],"Kprqq":"2022-05-16","Kprqz":"2022-06-01","Tfrqq":"2022-05-16","Tfrqz":"2022-06-01","DtBz":"N","PageNumber":1,"PageSize":10}`)
	req, err := http.NewRequest("POST", "https://dppt.neimenggu.chinatax.gov.cn:8443/szzhzz/qlfpcx/v1/queryFpjcxx?t=1653963547069", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "SSOTicket=092f3347-e169-464b-a99c-57b3fc23eef4; user=18d848d370c44e419fe6399b2f1fff98@; dzfp-ssotoken=db172a79613e4d79b8c1c4dfe2d8b9e6; SSO_SECURITY_CHECK_TOKEN=08fb227c3c184444aa26a98431c3e8a2; x_host_key=18117cc7df7-2fda0ec402f2a4bc5a92afb32597aa255a3cf333")
	req.Header.Set("Origin", "https://dppt.neimenggu.chinatax.gov.cn:8443")
	req.Header.Set("Referer", "https://dppt.neimenggu.chinatax.gov.cn:8443/invoice-query/invoice-query")
	req.Header.Set("SSO_SECURITY_CHECK_TOKEN", "08fb227c3c184444aa26a98431c3e8a2")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.67 Safari/537.36")
	req.Header.Set("nsrsbh", "91150103MA13QAX51F")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Google Chrome";v="101"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}

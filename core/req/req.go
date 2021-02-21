package req

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/spf13/cast"
)

type Param map[string]interface{}
type FormParam Param
type QueryParam map[string]interface{}
type Header map[string]string
type Proxy string

func (param Param) Encode() string {
	if param == nil {
		return ""
	}
	var buf strings.Builder
	for k, v := range param {
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(cast.ToString(v))
	}
	return buf.String()
}

type Resp struct {
	Data   []byte
	Header map[string][]string
}

func Do(method, rawURL string, vs ...interface{}) (result *Resp, err error) {
	if rawURL == "" {
		return nil, errors.New("url 为 空")
	}
	var formParam FormParam
	var queryParam QueryParam
	var header Header
	var proxy Proxy

	for _, v := range vs {
		switch vv := v.(type) {
		case FormParam:
			formParam = vv
		case QueryParam:
			queryParam = vv
		case Header:
			header = vv
		case Proxy:
			proxy = vv
		}
	}

	if len(queryParam) != 0 {
		queryStr := Param(queryParam).Encode()
		rawURL = rawURL + "?" + queryStr
	}

	url, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	req := &http.Request{
		Method: method,
		URL:    url,
		Header: make(http.Header),
		Body:   ioutil.NopCloser(strings.NewReader(Param(formParam).Encode())),
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	if proxy != "" {
		proxyURL, err := url.Parse(string(proxy))
		if err != nil {
			return nil, err
		}
		netTransport := &http.Transport{
			Proxy:                 http.ProxyURL(proxyURL),
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * 5,
		}
		client = &http.Client{
			Transport: netTransport,
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result = &Resp{
		Data:   body,
		Header: resp.Header,
	}
	return result, nil
}

func Get(rawURL string, vs ...interface{}) (result *Resp, err error) {
	return Do("GET", rawURL, vs...)
}

func Post(rawURL string, vs ...interface{}) (result *Resp, err error) {
	return Do("POST", rawURL, vs...)
}

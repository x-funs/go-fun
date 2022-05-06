package fun

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	DefaultTimeOut   = 5
	DefaultUserAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36"
)

// HttpGetBody Http Get，参数为请求地址，超时时间，请求头 map[string]string，返回值为请求内容，错误信息
func HttpGetBody(urlStr string, timeout int, headers map[string]string) (string, error) {
	if timeout == 0 {
		timeout = DefaultTimeOut
	}

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return "", err
	}

	// 处理请求头
	if headers == nil || len(headers) == 0 {
		headers = make(map[string]string)
		headers["User-Agent"] = DefaultUserAgent
	} else {
		if _, exist := headers["User-Agent"]; !exist {
			headers["User-Agent"] = DefaultUserAgent
		}
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	return string(body), nil
}

// HttpPostBody Http Post Form，参数为请求地址，Form 数据map[string]string，超时时间，请求头 map[string]string，返回值为请求内容，错误信息
func HttpPostBody(urlStr string, posts map[string]string, timeout int, headers map[string]string) (string, error) {
	if timeout == 0 {
		timeout = DefaultTimeOut
	}

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	data := url.Values{}
	if posts != nil && len(posts) > 0 {
		for k, v := range posts {
			data.Set(k, v)
		}
	}

	// 处理 Form
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	// 处理请求头
	if headers == nil || len(headers) == 0 {
		headers = make(map[string]string)
		headers["User-Agent"] = DefaultUserAgent
	} else {
		if _, exist := headers["User-Agent"]; !exist {
			headers["User-Agent"] = DefaultUserAgent
		}
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	return string(body), nil
}

// HttpPostJsonBody Http Post Json，参数为请求地址，Json数据string，超时时间，请求头 map[string]string，返回值为请求内容，错误信息
func HttpPostJsonBody(urlStr string, json string, timeout int, headers map[string]string) (string, error) {
	if timeout == 0 {
		timeout = DefaultTimeOut
	}

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// 处理 Json
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(json))
	if err != nil {
		return "", err
	}

	// 处理请求头
	if headers == nil || len(headers) == 0 {
		headers = make(map[string]string)
		headers["User-Agent"] = DefaultUserAgent
	} else {
		if _, exist := headers["User-Agent"]; !exist {
			headers["User-Agent"] = DefaultUserAgent
		}
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	return string(body), nil
}

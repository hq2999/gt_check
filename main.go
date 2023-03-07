package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

var hc *http.Client

func main() {

	var wg = sync.WaitGroup{}

	var ip_list []string

	hc = &http.Client{Timeout: 5 * time.Second}

	if strings.HasPrefix(os.Args[1], "http") {

		request, err := http.NewRequest("GET", os.Args[1], nil)
		request.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
		request.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")

		if err != nil {
			panic(err)
		}

		resp, err := hc.Do(request)

		if err != nil {
			panic(err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		resp.Body.Close()

		ip_list = strings.Split(string(body), "\n")

	} else if os.Args[1] == "-f" {
		// file
		f, err := ioutil.ReadFile(os.Args[2])
		if err != nil {
			panic(err)
		}
		ip_list = strings.Split(string(f), "\n")
	} else {
		ip_list = os.Args[1:]
	}

	for i := 0; i < len(ip_list); i++ {
		ip := strings.ReplaceAll(ip_list[i], "\r", "")
		if len(ip) > 0 {
			go do_check(ip, &wg)
			wg.Add(1)
		}
	}

	wg.Wait()
}

func do_check(addr string, wg *sync.WaitGroup) {

	url_ := fmt.Sprintf("http://%s/translate_a/single?client=gtx&sl=zh-CN&tl=en&dt=t&q=%s", addr, url.QueryEscape("你好"))

	request, err := http.NewRequest("GET", url_, nil)
	request.Host = "translate.googleapis.com"
	request.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")

	if err != nil {
		panic(err)
	}

	begin := time.Now()
	resp, err := hc.Do(request)

	if err != nil {
		wg.Done()
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		wg.Done()
		return
	}

	delay := time.Since(begin)

	resp.Body.Close()
	if strings.Contains(string(body), "Hello") {
		fmt.Printf("%-20s %s\n", addr, delay)
	}

	wg.Done()
}

package scrapy

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Spider struct {
	DataType string
}

// 百度测试
func (spider Spider) Baidu() []map[string]interface{} {
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	url := "https://www.baidu.com"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("抓取" + spider.DataType + "失败1")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36`)

	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + spider.DataType + "失败2")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(doc)
	doc.Find(".s-hotsearch-content .hotsearch-item").Each(func(i int, s *goquery.Selection) {
		content := s.Find(".title-content-title").Text()
		fmt.Printf("%d: %s\n", i, content)
	})
	var allData []map[string]interface{}
	return allData
}
func (spider Spider) BaiduB2b() []map[string]interface{} {
	var url string = "https://b2b.baidu.com/c/a?ajax=1&csrf_token=5fd0a18d629d86d617f0d5964f8c5c49&logid=3543431996693398577&fid=0%2C1657175877484&_=1657176095974&q=%E5%85%AC%E5%8F%B8&p=4&s=15&o=1&f=[{%22search_location%22:[%22%E9%83%91%E5%B7%9E%22]},{%22member_type%22:[0]}]"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("失败")
		return []map[string]interface{}{}
	}
	body, _ := ioutil.ReadAll(res.Body)
	data, err2 := simplejson.NewJson(body)
	if err2 != nil {
		fmt.Println("解析失败")
		return []map[string]interface{}{}
	}
	test := data.Get("data").Get("entList")
	fmt.Println(test)

	var allData []map[string]interface{}
	return allData
}

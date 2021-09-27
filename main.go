package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	// 创建收集器
	c := colly.NewCollector(
		// 设置UserAgent
		colly.UserAgent("Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Mobile Safari/537.36"),
		// 设置响应体中body的字节数限制、0为不限制、避坑：当下载文件时、不设置此项，下载的文件会以10兆为限制
		colly.MaxBodySize(0),
	)

	// 请求回调函数
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))

	})

	url := "https://www.sohu.com/"
	err := c.Visit(url)
	if err != nil {
		fmt.Println("err visit", err)
	}
}
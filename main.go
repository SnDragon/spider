package main

import (
	"fmt"
	"github.com/SnDragon/spider/collect"
	"regexp"
)

// .通配符无法匹配换行符，而 HTML 文本中会经常出现换行符,所以用\s\S
var headerRe = regexp.MustCompile("<div class=\"small_imgposition__PYVLm\">[\\s\\S]*?<h2>([\\s\\S]*?)</h2>")

func main() {
	//url := "https://www.thepaper.cn/"
	url := "https://book.douban.com/subject/1007305"
	//var f collect.Fetcher = &collect.BaseFetcher{}
	var f collect.Fetcher = &collect.BrowserFetcher{}
	rsp, err := f.Get(url)
	if err != nil {
		fmt.Printf("read body err: %+v\n", err)
		return
	}
	fmt.Printf("rsp body: %v", string(rsp))
	// 1. 正则表达式
	//matches := headerRe.FindAllSubmatch(bytes, -1)
	//for _, match := range matches {
	//	fmt.Printf("match: %v\n", string(match[1]))
	//}
	// 2. 使用xpath获取
	//doc, _ := htmlquery.Parse(bytes.NewReader(rsp))
	//nodes := htmlquery.Find(doc, `//div[@class="small_cardcontent__BTALp"]//h2`)
	//for _, node := range nodes {
	//	fmt.Printf("match: %v\n", node.FirstChild.Data)
	//}
	// 3. css表达式
	//doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(rsp))
	//doc.Find(".small_cardcontent__BTALp h2").Each(func(i int, s *goquery.Selection) {
	//	content := s.Text()
	//	fmt.Printf("match %s\n", content)
	//})
}

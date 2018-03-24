package main

import (
	"github.com/gocolly/colly"
	"fmt"
	"strings"
	"os"
	"regexp"
)

func main() {

	hrefMap := getAllSpiderUrl()

	f, err := os.OpenFile("./aaaa.txt", os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("can not open this file , err :" + err.Error())
		return
	}

	for key, value := range hrefMap{
		result1 := spiderByStringReplace("https://www.piaotian.com/html/3/3317/" + value)
		if result1 != "" {
			if key == 0 {
				result1 = "剑道独神\n" + result1 + "\n"
			}
			f.WriteString(result1)
			fmt.Printf("第%d抓取完成,数据正确\n", key)

		}else {
			fmt.Printf("第%d抓取完成，数据错误\n", key)
		}
	}
}

func getAllSpiderUrl() []string {
	var result []string
	url := "https://www.piaotian.com/html/3/3317/"
	c := colly.NewCollector()
	c.OnHTML("div", func(e *colly.HTMLElement)  {
		e.ForEach("a", func(i int, element *colly.HTMLElement) {
			var match bool
			//fmt.Println(element.Attr("href"))
			//e.Request.Visit(element.Attr("href"))
			href := element.Attr("href")
			match, _ = regexp.MatchString("^[0-9]+.html$", href)
			if match == true {
				result = append(result, href)
			}
		})

	})
	//return []string{}
	c.Visit(url)
	//return result
	return result
}

func spiderByStringReplace(url string) string {
	var link string
	c := colly.NewCollector()
	c.OnHTML("body", func(e *colly.HTMLElement)  {
		link = e.Text

		if !strings.Contains(link,"剑道独神") {
			link = ""
			return
		}

		if strings.Contains(link, "访问错误") {
			link = ""
			return
		}
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		child := e.ChildText("div")
		childSplit := strings.Split(child,"\n",)

		for _, value := range childSplit  {
			link = strings.Replace(link,value,"",-1)
		}
		link = strings.Replace(link,"Gundong();","",-1)
		link = strings.Replace(link,"GetFont();","",-1)
		link = strings.Replace(link,"{飘天文学www.piaotian.com感谢各位书友的支持，您的支持就是我们最大的动力}","",-1)
		link = strings.Replace(link,"推荐本书上一章目 录下一章加入书签","",-1)
		link = strings.Replace(link,"温馨提示：方向键左右(← →)前后翻页，上下(↑ ↓)上下滚用， 回车键:返回目录","",-1)
		link = strings.Replace(link,"推荐阅读：","",-1)
		link = strings.Replace(link,"返回顶部","",-1)
		link = strings.Replace(link,"我的藏书架","",-1)
		link = strings.Replace(link,"重要声明：小说“”所有的文字、目录、评论、图片等，均由网友发表或上传并维护或来自搜索引擎结果，属个人行为，与本站立场无关。阅读更多小说最新章节请返回飘天文学网首页，小说阅读网永久地址：www.piaotian.netCopyright © 2012-2013 飘天文学-飘越天空的小说阅读网  All rights reserved.","",-1)
		link = strings.Replace(link,"&nbsp","",-1)
		link = strings.Replace(link,"Process","",-1)
		link = strings.TrimSpace(link)
	})
	//685240
	//url = "https://www.piaotian.com/html/3/3317/685240.html"
	c.Visit(url)
	return link
}

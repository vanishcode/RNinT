// Package fetcher 目标网站的章节列表
package fetcher

import (
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"github.com/vanishcode/RNinT/utils"
)

// Chapter 章节抽象
type Chapter struct {
	Idx int
	Title string
	Link string
}

// ParseChapterList 章节列表请求、解析
func ParseChapterList(url string) ([]Chapter){
	var result []Chapter
	var index = 0
	u := utils.EncodeURIComponent(url)
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", utils.BasicEngineURL + u,nil)
	if err != nil {
	  log.Fatal(err)
	}
	reqest.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	res, err := client.Do(reqest)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
	  log.Fatal(err)
	}
	doc.Find("dd").Each(func(i int, s *goquery.Selection) {
	  title := s.Find("a").Text()
	  href, _ := s.Find("a").Attr("href")

	  var item Chapter
	  item.Idx = index
	  item.Title = title
	  item.Link = href
	  result = append(result,item)

	  index++
	})
	return result;
}

// Package fetcher 搜索 + 结果列表
// e.g. https://www.owllook.net/search?wd=龙族
package fetcher

import (
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"github.com/vanishcode/RNinT/utils"
)

// Item 结果页每一项的结构
type Item struct {
	Keyword string
	URL string
}

// Search 搜索、解析结果页功能
func Search(name string) ([]Item){
	var result []Item
	encodeName := utils.EncodeURIComponent(name)
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", utils.SearchURL+encodeName,nil)
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
	doc.Find(".result_item > li").Each(func(i int, s *goquery.Selection) {
	  keyword := s.Find("a").Text()
	  href, _ := s.Find("a").Attr("href")
	  flag := strings.Trim(s.Find(".label-primary").Text()," ")
	  if flag == "已解析" ||  flag == "推荐源已解析"{
		var item Item
		item.Keyword = keyword
		item.URL = href
		result = append(result,item)
	  }
	})
	return result;
  }

// Package fetcher 要阅读的章的内容
package fetcher

import (
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"github.com/vanishcode/RNinT/utils"
)

// GetContent 获取章节内容，解析规则一定
func GetContent(url string) string {
	u := utils.EncodeURIComponent(url)
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", utils.BasicEngineURL + "/" + u,nil)
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
	content := doc.Find("#content").Text()
	return content
}
package parser

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/herzorf/gun/engine"
	"log"
	"strings"
)

func MHYArticleListParse(content []byte) engine.ParserResult {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(content))
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".mhy-article-card").Each(func(i int, selection *goquery.Selection) {
		userName := selection.Find(".mhy-account-title__name").Text()
		article := selection.Find(".mhy-router-link.mhy-article-card__link")
		topics := selection.Find(".mhy-article-card__topics span")
		info := selection.Find(".mhy-article-card__data span")
		title := selection.Find(".mhy-article-card__title h3")
		val, _ := article.Attr("href")
		date := selection.Find(".mhy-article-card__info").First().Text()
		fmt.Println(strings.TrimSpace(userName))
		fmt.Println(strings.TrimSpace(date))
		fmt.Println(title.Text())

		fmt.Println(" https://www.miyoushe.com/" + val)
		var t []string
		topics.Each(func(i int, selection *goquery.Selection) {
			//fmt.Printf("标签：%d: %s \n", i, selection.Text())
			t = append(t, selection.Text())
		})
		fmt.Printf("%v\n", t)
		fmt.Println(info.Text())
	})
	//compile := regexp.MustCompile(cityListReg)
	//submatch := compile.FindAllStringSubmatch(string(content), -1)
	result := engine.ParserResult{}
	//for _, value := range submatch {
	//	result.Item = append(result.Item, value[2])
	//	result.Requests = append(result.Requests, engine.Request{
	//		Url:        value[1],
	//		ParserFunc: engine.NilParser,
	//	})
	//}
	return result
}

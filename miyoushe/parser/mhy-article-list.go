package parser

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/herzorf/gun/engine"
	"strings"
)

func MHYArticleListParse(content []byte) engine.ParserResult {
	result := engine.ParserResult{}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(content))
	if err != nil {
		panic(err)
	}
	doc.Find(".mhy-article-card").Each(func(i int, selection *goquery.Selection) {
		userName := selection.Find(".mhy-account-title__name").Text()
		article := selection.Find(".mhy-router-link.mhy-article-card__link")
		topics := selection.Find(".mhy-article-card__topics span")
		info := selection.Find(".mhy-article-card__data span")
		title := selection.Find(".mhy-article-card__title h3")
		link, _ := article.Attr("href")
		date := selection.Find(".mhy-article-card__info").First().Text()
		var t []string
		topics.Each(func(i int, selection *goquery.Selection) {
			t = append(t, selection.Text())
		})
		var in map[string]string
		in = make(map[string]string)
		key := [3]string{"浏览", "评论", "点赞"}
		info.Each(func(i int, selection *goquery.Selection) {
			in[key[i]] = selection.Text()
		})
		var articleCard = ArticleCard{
			Username: strings.TrimSpace(userName),
			Date:     strings.TrimSpace(date),
			Title:    title.Text(),
			Link:     " https://www.miyoushe.com" + link,
			Topics:   t,
			Info:     in,
		}
		result.Requests = append(result.Requests, engine.Request{
			Url:        articleCard.Link,
			Data:       articleCard,
			ParserFunc: engine.NilParser,
		})
		fmt.Printf("%+v\n", articleCard)
	})
	return result
}

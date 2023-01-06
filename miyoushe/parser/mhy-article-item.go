package parser

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/herzorf/gun/engine"
)

func MHYArticleItemParse(request engine.Request, content []byte) engine.ParserResult {
	var dataSrc []string
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(content))
	if err != nil {
		panic(err)
	}
	doc.Find(".mhy-article-page__content .mhy-img-article img").Each(func(i int, selection *goquery.Selection) {
		src, exist := selection.Attr("src")
		if exist {
			dataSrc = append(dataSrc, src)
		}
	})
	doc.Find(".mhy-video-article-container .mhy-video-player__video video").Each(func(i int, selection *goquery.Selection) {
		src, exist := selection.Attr("src")
		if exist {
			dataSrc = append(dataSrc, src)
		}
	})
	request.Data.DataSrc = dataSrc

	fmt.Printf("%+v", request)
	return engine.ParserResult{}
}

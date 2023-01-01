package fetcher

import (
	"context"
	"github.com/chromedp/chromedp"
)

func Fetcher(url string) ([]byte, error) {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list.
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.OuterHTML(`document.querySelector(".mhy-article-list")`, &res, chromedp.ByJSPath),
	)
	if err != nil {
		return []byte(nil), err
	}
	return []byte(res), nil

}

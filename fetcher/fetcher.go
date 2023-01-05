package fetcher

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
)

func Fetcher(url string) ([]byte, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.OuterHTML(fmt.Sprintf(`document.querySelector(".mhy-layout__main")`), &res, chromedp.ByJSPath),
	)
	if err != nil {
		return []byte(nil), err
	}
	return []byte(res), nil

}

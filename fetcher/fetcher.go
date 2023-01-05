package fetcher

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"time"
)

func Fetcher(url string) ([]byte, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
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

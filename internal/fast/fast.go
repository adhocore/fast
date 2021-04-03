package fast

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/cdproto/emulation"
)

func Run() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	start := time.Now()

	var (
		up     string
		dw     string
		upunit string
		dwunit string
	)

	err := chromedp.Run(ctx,
		emulation.SetUserAgentOverride(`chromedp/chromedp v0.6.10`),
		chromedp.Navigate(`https://fast.com`),
		chromedp.ScrollIntoView(`footer`),
		chromedp.WaitVisible(`#speed-value.succeeded`),
		chromedp.Text(`#speed-value.succeeded`, &dw, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#speed-units.succeeded`, &dwunit, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Click(`#show-more-details-link`),
		chromedp.WaitVisible(`#upload-value.succeeded`),
		chromedp.Text(`#upload-value.succeeded`, &up, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#upload-units.succeeded`, &upunit, chromedp.NodeVisible, chromedp.ByQuery),
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\033[36mdownload speed:\033[m \033[32m%s\033[m %s\n", dw, dwunit)
	fmt.Printf("\033[36mupload speed:\033[m \033[31m%s\033[m %s\n", up, upunit)
	fmt.Printf("\n")
	fmt.Printf("\033[36m> took: \033[33m%f\033[m secs\n", time.Since(start).Seconds())
}

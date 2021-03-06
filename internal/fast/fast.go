package fast

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
)

// Fast represents measurement structure
type Fast struct {
	Up       string
	Down     string
	UpUnit   string
	DownUnit string
}

// Measure does the main job.
// It returns *Fast and error
func Measure(noUp bool) (*Fast, error) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 180*time.Second)
	defer cancel()

	fast := new(Fast)
	cmds := []chromedp.Action{
		emulation.SetUserAgentOverride(`chromedp/chromedp v0.6.10`),
		chromedp.Navigate(`https://fast.com`),
		chromedp.ScrollIntoView(`footer`),
		chromedp.WaitVisible(`#speed-value.succeeded`),
		chromedp.Text(`#speed-value.succeeded`, &fast.Down, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#speed-units.succeeded`, &fast.DownUnit, chromedp.NodeVisible, chromedp.ByQuery),
	}

	if !noUp {
		cmds = append(cmds, chromedp.Click(`#show-more-details-link`),
			chromedp.WaitVisible(`#upload-value.succeeded`),
			chromedp.Text(`#upload-value.succeeded`, &fast.Up, chromedp.NodeVisible, chromedp.ByQuery),
			chromedp.Text(`#upload-units.succeeded`, &fast.UpUnit, chromedp.NodeVisible, chromedp.ByQuery),
		)
	}

	err := chromedp.Run(ctx, cmds...)

	return fast, err
}

// Run is the ready to use API.
// For customization call Measure().
func Run(noUp bool) {
	start := time.Now()

	fast, err := Measure(noUp)
	if err != nil {
		log.Fatal(err)
	}

	Out(fast, start)
}

// Out prints the output to terminal
func Out(fast *Fast, start time.Time) {
	hasUp := fast.Up != "" && fast.UpUnit != ""

	// No color for windows
	if os.PathSeparator == '\\' {
		fmt.Printf("download speed: %s %s\n", fast.Down, fast.DownUnit)
		if hasUp {
			fmt.Printf("upload speed: %s %s\n", fast.Up, fast.UpUnit)
		}
		fmt.Printf("\n")
		fmt.Printf("> took: %f secs\n", time.Since(start).Seconds())

		return
	}

	fmt.Printf("\033[36mdownload speed:\033[m \033[32m%s\033[m %s\n", fast.Down, fast.DownUnit)
	if hasUp {
		fmt.Printf("\033[36mupload speed:\033[m \033[31m%s\033[m %s\n", fast.Up, fast.UpUnit)
	}
	fmt.Printf("\n")
	fmt.Printf("\033[36m> took: \033[33m%f\033[m secs\n", time.Since(start).Seconds())
}

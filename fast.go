package fast

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
)

// Browsers is list of non default browsers by OS
var Browsers = map[string][]string{
	"windows": {"", `C:\Program Files\BraveSoftware\Brave-Browser\Application\brave.exe`},
	"darwin":  {"", `/Applications/Brave Browser.app/Contents/MacOS/BraveBrowser`},
	"*":       {"", `brave`, `brave-browser`},
}

// NoColor tells if color should be used
var NoColor = os.Getenv("NO_COLOR") != "" || os.Getenv("TERM") == "dumb"

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
	fast := new(Fast)
	cmds := []chromedp.Action{
		emulation.SetUserAgentOverride(`chromedp/chromedp v0.8.6`),
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

	browsers := Browsers["*"]
	if paths, ok := Browsers[runtime.GOOS]; ok {
		browsers = paths
	}

	var err error
	for _, browser := range browsers {
		err = doMeasure(browser, cmds...)
		if err == nil || !errors.Is(err, exec.ErrNotFound) {
			break
		}
	}

	return fast, err
}

func doMeasure(browser string, cmds ...chromedp.Action) error {
	opts := chromedp.DefaultExecAllocatorOptions[:]
	if browser != "" {
		opts = append(opts, chromedp.ExecPath(browser))
	}

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 180*time.Second)
	defer cancel()

	return chromedp.Run(ctx, cmds...)
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

	out := fmt.Sprintf("\033[36mdownload speed:\033[m \033[32m%s\033[m %s\n", fast.Down, fast.DownUnit)
	if hasUp {
		out += fmt.Sprintf("\033[36mupload speed:  \033[m \033[31m%s\033[m %s\n", fast.Up, fast.UpUnit)
	}
	out += fmt.Sprintf("\n\033[36m> took: \033[33m%f\033[m secs\n", time.Since(start).Seconds())

	if NoColor {
		out = strings.ReplaceAll(strings.ReplaceAll(out, "\033[36m", ""), "\033[m", "")
		out = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(out, "\033[33m", ""), "\033[32m", ""), "\033[31m", "")
	}

	fmt.Print(out)
}

package main

import (
	"flag"
	"log"
	"sync"
	"time"

	"github.com/adhocore/chin"
	"github.com/adhocore/fast"
)

var noUp bool

func init() {
	flag.BoolVar(&noUp, "noup", false, "Do not show upload speed (shows only download speed)")
	flag.Parse()
}

func main() {
	var wg sync.WaitGroup

	s := chin.New().WithWait(&wg)
	go s.Start()

	wg.Add(1)
	go doFast(s, &wg, noUp)

	wg.Wait()
}

func doFast(s *chin.Chin, wg *sync.WaitGroup, noUp bool) {
	defer wg.Done()

	start := time.Now()
	res, err := fast.Measure(noUp)

	if err != nil {
		log.Fatalf("error measuring speed: %v", err)
	}

	s.Stop()
	fast.Out(res, start)
}

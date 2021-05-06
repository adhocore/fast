package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/adhocore/fast/internal/fast"
)

func main() {
	var noUp bool
	var wg sync.WaitGroup

	flag.BoolVar(&noUp, "noup", false, "Do not show upload speed (shows only download speed)")
	flag.Parse()
	wg.Add(1)

	ch := make(chan bool)

	go doSpin(ch)
	go doFast(ch, &wg, noUp)

	wg.Wait()
}

func doSpin(ch chan bool) {
	chars := []string{"+", "\\", "|", "/", "-", "+", "\\", "|", "/", "-"}

	for {
	outer:
		select {
		case _, ok := <-ch:
			if ok {
				fmt.Print("\010")
				break outer
			}
		default:
			for _, c := range chars {
				fmt.Print(c, "\010")
				time.Sleep(50 * time.Millisecond)
			}
		}
	}
}

func doFast(ch chan bool, wg *sync.WaitGroup, noUp bool) {
	defer wg.Done()

	start := time.Now()
	res, err := fast.Measure(noUp)
	ch <- true

	if err != nil {
		log.Fatalf("error measuring speed: %v", err)
	}

	fast.Out(res, start)
}

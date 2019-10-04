package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	format, duration, font := parseArgs()
	displayDuration := time.Duration(duration) * time.Second

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	canceled := false
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	go func() {
		<-sc
		cancel()
		canceled = true
	}()
	goFmtStr := formatDate(format)

	cnt := 0
	for now := range now(ctx, displayDuration) {
		print("\033[H\033[2J")
		nowStr := now.Format(goFmtStr)
		for _, str := range strings.Split(nowStr, "\n") {
			figure.NewFigure(str, font, true).Print()
		}
		cnt++
	}
	if canceled {
		fmt.Println("canceled!")
	}
}

func now(ctx context.Context, duration time.Duration) <-chan time.Time {
	ch := make(chan time.Time)
	go func(ch chan<- time.Time) {
		defer close(ch)
	loop:
		for now := range time.Tick(duration) {
			select {
			case <-ctx.Done():
				break loop
			default:
			}
			ch <- now
		}
	}(ch)
	return ch
}

func formatDate(format string) string {
	var matchTypes = []struct {
		dateCmdFmt string
		goFmt      string
	}{
		{
			dateCmdFmt: "%Y",
			goFmt:      "2006",
		},
		{
			dateCmdFmt: "%m",
			goFmt:      "01",
		},
		{
			dateCmdFmt: "%d",
			goFmt:      "02",
		},
		{
			dateCmdFmt: "%H",
			goFmt:      "15",
		},
		{
			dateCmdFmt: "%M",
			goFmt:      "04",
		},
		{
			dateCmdFmt: "%S",
			goFmt:      "05",
		},
	}
	goFmtStr := format
	for _, m := range matchTypes {
		goFmtStr = strings.Replace(goFmtStr, m.dateCmdFmt, m.goFmt, -1)
	}
	return goFmtStr
}

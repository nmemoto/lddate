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
	format, duration, font, position := parseArgs()
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
		var artStr artStr
		for _, str := range strings.Split(nowStr, "\n") {
			artStr = append(artStr, figure.NewFigure(str, font, true).Slicify())
		}

		switch position {
		case "left":
			artStr.Print()
		case "center":
			artStr.CenterPrint()
		case "right":
			artStr.RightPrint()
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

type artStr [][]string

func (str artStr) Print() {
	for _, rowArtStr := range str {
		for _, row := range rowArtStr {
			fmt.Println(row)
		}
	}
}

func (str artStr) CenterPrint() {
	maxWidth := 0
	for _, rowArtStr := range str {
		if maxWidth < len(rowArtStr[0]) {
			maxWidth = len(rowArtStr[0])
		}
	}
	for i, rowArtStr := range str {
		if maxWidth == len(rowArtStr[0]) {
			continue
		}
		diff := (maxWidth - len(rowArtStr[0])) / 2
		for k, rowStr := range rowArtStr {
			newRowStr := strings.Repeat(" ", diff) + rowStr
			str[i][k] = newRowStr
		}
	}
	str.Print()
}

func (str artStr) RightPrint() {
	maxWidth := 0
	for _, rowArtStr := range str {
		if maxWidth < len(rowArtStr[0]) {
			maxWidth = len(rowArtStr[0])
		}
	}
	for i, rowArtStr := range str {
		if maxWidth == len(rowArtStr[0]) {
			continue
		}
		diff := maxWidth - len(rowArtStr[0])
		for k, rowStr := range rowArtStr {
			newRowStr := strings.Repeat(" ", diff) + rowStr
			str[i][k] = newRowStr
		}
	}
	str.Print()
}
